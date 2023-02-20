package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string         `gorm:"primaryKey;type:varchar(255)" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	Email     string         `json:"email" gorm:"not null; type:varchar(255)"`
	Password  string         `json:"password" gorm:"not null; type:varchar(255)"`
	Notes     []Note         `json:"notes" gorm:"foreignKey:user_id"`
}
