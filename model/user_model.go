package model

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	ID         int64  `gorm:"primaryKey;autoIncrement"`
	Identity   string `gorm:"not null;"`
	Email      string `gorm:"not null;"`
	Password   string `gorm:"not null;"`
	CreateTime time.Time
	UpdateTime time.Time
}

func (m *UserBasic) TableName() string {
	return "user_basic"
}

func CreateUser(ub *UserBasic) error {
	db := GetDB()
	if err := db.Create(ub).Error; err != nil {
		return err
	}
	return nil
}

func GetUserByEmailAndPassword(email, password string) (*UserBasic, error) {
	db := GetDB()
	var ub UserBasic
	if err := db.Where("email = ? AND password = ?", email, password).First(&ub).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &ub, nil
}

func GetUserByEmail(email string) (*UserBasic, error) {
	db := GetDB()
	var ub UserBasic
	if err := db.Where("email = ?", email).First(&ub).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &ub, nil
}

func GetUserByIdentity(identity string) (*UserBasic, error) {
	db := GetDB()
	var ub UserBasic
	if err := db.Where("identity = ?", identity).First(&ub).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &ub, nil
}
