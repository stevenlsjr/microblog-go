package models

import (
	"gorm.io/gorm"
	"time"
)

type UuidModel struct {
	ID        string         `gorm:"primarykey;default:gen_random_uuid()" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

func NewUuidModel(ID string) *UuidModel {
	now := time.Now()
	model := UuidModel{ID: ID, CreatedAt: now, UpdatedAt: now}
	_ = model.DeletedAt.Scan(nil)
	return &model
}

type BlogUser struct {
	UuidModel
	Username string `gorm:"unique;index" json:"username"`
}

func AllModels() []interface{} {
	return []interface{}{
		&BlogUser{},
	}
}
