package controller

import (
	"net/http"
	"time"

	"gin-calender/model"
	"gin-calender/utils"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	var ub model.UserBasic
	ctx.Bind(&ub)
	email := ub.Email
	password := ub.Password

	//数据验证
	if email == "" {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 400,
			"msg":  "邮箱为空",
		})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 400,
			"msg":  "密码不能少于6位",
		})
		return
	}

	user, err := model.GetUserByEmail(email)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 500,
			"msg":  "内部错误" + err.Error(),
		})
		return
	}
	if user != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 400,
			"msg":  "用户已存在",
		})
		return
	}
	user = &model.UserBasic{
		Identity:   utils.GetUUID(),
		Email:      email,
		Password:   password,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	err = model.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 500,
			"msg":  "用户创建失败" + err.Error(),
		})
		return
	}

	//返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "注册成功",
	})
}

func Login(ctx *gin.Context) {
	var ub model.UserBasic
	ctx.Bind(&ub)
	email := ub.Email
	password := ub.Password

	//数据验证
	if email == "" {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 400,
			"msg":  "邮箱为空",
		})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 400,
			"msg":  "密码不能少于6位",
		})
		return
	}

	user, err := model.GetUserByEmail(email)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 500,
			"msg":  "内部错误" + err.Error(),
		})
		return
	}
	if user == nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 400,
			"msg":  "用户不存在",
		})
		return
	}
	if user.Password != password {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 400,
			"msg":  "密码错误",
		})
		return
	}
	user, err = model.GetUserByEmailAndPassword(email, password)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 500,
			"msg":  "内部错误",
		})
		return
	}
	token, err := utils.GenerateToken(user.Identity, user.Email)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 500,
			"msg":  "内部错误" + err.Error(),
		})
		return
	}

	//返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"data": gin.H{
			"token": token,
		},
	})
}
