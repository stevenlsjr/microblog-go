package models

import (
	"gorm.io/gorm"
	"time"
)

type UuidModel struct {
	ID        string         `gorm:"primarykey;default:gen_random_uuid()" json:"id" faker:"uuid_hyphenated"`
	CreatedAt time.Time      `json:"createdAt" faker:"-"`
	UpdatedAt time.Time      `json:"updatedAt" faker:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt" faker:"-"`
}

func NewUuidModel(ID string) *UuidModel {
	now := time.Now()
	model := UuidModel{ID: ID, CreatedAt: now, UpdatedAt: now}
	_ = model.DeletedAt.Scan(nil)
	return &model
}

func AllModels() []interface{} {
	return []interface{}{
		&User{},
		&Message{},
	}
}
