package controller

import (
	"fmt"
	"net/http"
	"time"

	"gin-calender/model"
	"gin-calender/utils"

	"github.com/gin-gonic/gin"
)

type ReminderInfo struct {
	Content      string `json:"content"`
	ReminderTime string `json:"reminder_time"`
}

// 创建提醒
func CreateReminder(ctx *gin.Context) {
	var info ReminderInfo
	err := ctx.Bind(&info)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数不足",
		})
		return
	}

	uc, err := utils.AnalyseToken(ctx.GetHeader("Authorization"))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  "Unauthorized",
		})
		return
	}

	loc, _ := time.LoadLocation("Asia/Shanghai")
	t, err := time.ParseInLocation("2006-01-02 15:04:05", info.ReminderTime, loc)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "时间格式错误",
		})
		return
	}

	reminderInfo := model.ReminderBasic{
		CreatorID:    uc.Identity,
		Content:      info.Content,
		ReminderTime: t,
		IsDeleted:    0,
		CreatedTime:  time.Now(),
		UpdatedTime:  time.Now(),
	}

	err = model.CreateReminderInfo(&reminderInfo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "创建失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "创建成功",
		"data": reminderInfo,
	})
}

// 获取本人的提醒列表
func GetReminderList(ctx *gin.Context) {
	uc, err := utils.AnalyseToken(ctx.GetHeader("Authorization"))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  "Unauthorized",
		})
		return
	}
	creatorID := uc.Identity

	reminderInfos, err := model.GetReminderInfosByCreatorID(creatorID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "获取失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": reminderInfos,
	})
}

// 更新提醒
func UpdateReminder(ctx *gin.Context) {
	reminderID := ctx.Param("id")
	var info ReminderInfo
	err := ctx.Bind(&info)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数不足",
		})
		return
	}

	uc, err := utils.AnalyseToken(ctx.GetHeader("Authorization"))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  "Unauthorized",
		})
		return
	}

	loc, _ := time.LoadLocation("Asia/Shanghai")
	t, err := time.ParseInLocation("2006-01-02 15:04:05", info.ReminderTime, loc)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "时间格式错误",
		})
		return
	}

	reminderInfo := model.ReminderBasic{
		CreatorID:    uc.Identity,
		Content:      info.Content,
		ReminderTime: t,
		IsDeleted:    0,
		UpdatedTime:  time.Now(),
	}

	err = model.UpdateReminderInfo(reminderID, &reminderInfo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "更新失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "更新成功",
		"data": reminderInfo,
	})
}

// 删除提醒
func DeleteReminder(ctx *gin.Context) {
	reminderID := ctx.Param("id")

	fmt.Printf("reminderID: %s\n", reminderID)

	uc, err := utils.AnalyseToken(ctx.GetHeader("Authorization"))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  "Unauthorized",
		})
		return
	}
	creatorID := uc.Identity

	err = model.DeleteReminderInfo(reminderID, creatorID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "删除失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除成功",
	})
}
