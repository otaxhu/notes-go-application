package models

import (
	"time"

	"gorm.io/gorm"
)

type Note struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	Title       string         `json:"title" gorm:"not null"`
	Description string         `json:"description" gorm:"not null"`
}
