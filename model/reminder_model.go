package model

import (
	"errors"
	"time"
)

type ReminderBasic struct {
	ID           int64     `gorm:"primaryKey;autoIncrement"`
	CreatorID    string    `gorm:"not null;"`
	Content      string    `gorm:"not null;"`
	ReminderTime time.Time `gorm:"not null;"`
	IsDeleted    int32     `gorm:"default:0;"`
	CreatedTime  time.Time
	UpdatedTime  time.Time
}

func (m *ReminderBasic) TableName() string {
	return "reminder"
}

func CreateReminderInfo(reminderInfo *ReminderBasic) error {
	db := GetDB()
	err := db.Create(reminderInfo).Error
	if err != nil {
		return err
	}

	return nil
}

func GetReminderInfosByCreatorID(creatorID string) ([]*ReminderBasic, error) {
	var reminderInfos []*ReminderBasic

	db := GetDB()
	err := db.Where("creator_id = ? AND is_deleted = 0", creatorID).Find(&reminderInfos).Error
	if err != nil {
		return nil, err
	}

	return reminderInfos, nil
}

func UpdateReminderInfo(id string, reminderInfo *ReminderBasic) error {
	db := GetDB()
	err := db.Model(reminderInfo).Where("id = ?", id).Updates(reminderInfo).Error
	if err != nil {
		return err
	}

	return nil
}

func DeleteReminderInfo(reminderID string, creatorID string) error {
	db := GetDB()
	result := db.Model(&ReminderBasic{}).Where("id = ? AND creator_id = ?", reminderID, creatorID).Update("is_deleted", 1)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("reminder not found or not owned by the user")
	}

	return nil
}
