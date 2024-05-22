package main

import (
	"gin-calender/config"
	"gin-calender/model"
	"gin-calender/router"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	// 初始化
	config.InitConfig()
	model.InitDB()
	r := gin.Default()
	router.SetupRouter(r)

	port := viper.GetString("server.port")
	panic(r.Run(":" + port))
}
