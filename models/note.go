package models

import (
	"time"

	"gorm.io/gorm"
)

type Note struct {
	ID          string         `gorm:"primaryKey;type:varchar(255)" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	Title       string         `json:"title" gorm:"not null;type:varchar(100)"`
	Description string         `json:"description" gorm:"not null;type:varchar(255)"`
	UserID      string         `json:"user_id" gorm:"not null"`
	User        User           `json:"user"`
}
