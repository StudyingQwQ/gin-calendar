package test

import (
	"fmt"
	"testing"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID         int64 `gorm:"primary_key"`
	Identity   string
	Email      string
	Password   string
	CreateTime time.Time
	UpdateTime time.Time
}

func TestUser(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/gin_calendar?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Errorf("failed to connect database, got error: %v", err)
		return
	}

	db.AutoMigrate(&User{})

	user1 := User{
		Identity:   "test1",
		Email:      "test1@qq.com",
		Password:   "123456",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}

	user2 := User{
		Identity:   "test2",
		Email:      "test2@qq.com",
		Password:   "123456",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}

	// 增加
	db.Create(&user1)
	db.Create(&user2)

	// 查询
	var result User
	db.First(&result, user1.ID)
	fmt.Printf("result: %+v\n", result)

	// 更新
	result.Email = "123@qq.com"
	db.Save(&result)
	fmt.Printf("result: %+v\n", result)

	// 删除
	db.Delete(&result)
}
