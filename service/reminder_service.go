package service

import (
	"context"
	"fmt"
	"gin-calender/model"
	"gin-calender/utils"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

func cacheRemindersToRedis(redisClient *redis.Client) {
	ctx := context.Background()

	// 从数据库中获取所有的日程
	reminders, err := model.GetAllReminderInfos()
	if err != nil {
		log.Printf("获取数据库失败: %v", err)
		return
	}

	// 将日程数据缓存到 Redis
	for _, reminder := range reminders {
		score := float64(reminder.ReminderTime.Unix())
		_, err := redisClient.ZAdd(ctx, "reminders", &redis.Z{
			Score:  score,
			Member: reminder.ID,
		}).Result()
		if err != nil {
			log.Printf("缓存失败: %v", err)
		}
	}

	log.Printf("缓存 %d 条日程至Redis", len(reminders))
}

func ProcessReminders(redisClient *redis.Client) {
	cacheRemindersToRedis(redisClient)

	ctx := context.Background()
	now := time.Now().Unix()

	// 从 Redis 中获取所有到期的日程
	reminderIDs, err := redisClient.ZRangeByScore(ctx, "reminders", &redis.ZRangeBy{Min: "0", Max: fmt.Sprintf("%d", now)}).Result()
	if err != nil {
		log.Printf("从Redis获取日程失败: %v", err)
		return
	}

	// 处理到期的日程
	for _, reminderID := range reminderIDs {
		// 从数据库中获取日程 详细信息
		var reminder model.ReminderBasic
		err := model.DB.First(&reminder, reminderID).Error
		if err != nil {
			log.Printf("从数据库获取日程失败: %v", err)
			continue
		}

		// 发送通知
		sendNotification(&reminder)
		// 更新缓存
		redisClient.ZRem(ctx, "reminders", reminderID)

		// 从 Redis 中删除已处理的日程
		_, err = redisClient.ZRemRangeByScore(ctx, "reminders", "0", fmt.Sprintf("%d", now)).Result()
		if err != nil {
			log.Printf("从缓存删除过期日程失败: %v", err)
		}

		// 从数据库中删除已处理的日程
		_ = model.DeleteReminderInfo(fmt.Sprintf("%d", reminder.ID), reminder.CreatorID)
	}
}

func sendNotification(reminder *model.ReminderBasic) {
	// 这里添加发送通知的逻辑
	log.Printf("用户: %s,你在时间点: %s需要去做: %s", reminder.CreatorID, reminder.ReminderTime, reminder.Content)
	ub, err := model.GetUserByIdentity(reminder.CreatorID)
	if err != nil {
		log.Printf("获取用户失败: %v", err)
		return
	}
	err = utils.MailReminder(ub.Email, reminder.Content)
	if err != nil {
		log.Printf("发送邮件失败: %v", err)
		return
	}
}
