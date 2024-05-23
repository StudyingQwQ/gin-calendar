package main

import (
	"gin-calender/config"
	"gin-calender/model"
	"gin-calender/router"
	"gin-calender/service"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	// 初始化
	config.InitConfig()
	model.InitDB()
	r := gin.Default()
	router.SetupRouter(r)

	model.InitDB()
	redisClient := model.RDB

	go func() {
		log.Printf("定时任务已启动")
		// 每 10 秒检查一次过期的 Reminder
		for {
			service.ProcessReminders(redisClient)
			time.Sleep(10 * time.Second)
		}
	}()

	port := viper.GetString("server.port")
	panic(r.Run(":" + port))
}
