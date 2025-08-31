package project

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string    `json:"name" gorm:"type:varchar(255);not null"`
	Description string    `json:"description" gorm:"type:text"`
	Archived    bool      `json:"archived" gorm:"default:false"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&Project{})
}
