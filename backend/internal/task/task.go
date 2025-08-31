package task

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID          uint      `json:"id" gorm:"primaryKey:autoIncrement"`
	Title       string    `json:"title"`
	Description string    `json:"description" gorm:"type:text"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `json:"status" gorm:"type:varchar(100);default:'pending'"`
	Priority    string    `json:"priority" gorm:"type:varchar(100);default:'low'"`
	ProjectID   uint      `json:"project_id"`
	Completed   bool      `json:"completed" gorm:"default:false"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&Task{})
}
