package router

import (
	"gin-calender/controller"
	"gin-calender/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) *gin.Engine {
	//用户模块
	r.POST("/register", controller.Register) //注册
	r.POST("/login", controller.Login)       //登录

	//提醒模块
	reminder := r.Group("/reminders")
	reminder.Use(middleware.JWTAuth())
	{
		reminder.POST("/", controller.CreateReminder)
		reminder.GET("/", controller.GetReminderList)
		reminder.PATCH("/:id", controller.UpdateReminder)
		reminder.DELETE("/:id", controller.DeleteReminder)
	}

	return r
}
