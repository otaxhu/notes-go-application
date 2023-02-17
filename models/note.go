package models

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	Title       string `json:"title" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
}
