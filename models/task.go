package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       string    `gorm:"not null" validate:"required"`
	Description string    `gorm:"not null" validate:"required"`
	Status      bool      `gorm:"not null" validate:"required"`
	UserId      uint      `json:"user_id"`
	CategoryId  uint      `json:"category_id"`
	User        *User     `gorm:"foreignKey:UserId"`
	Category    *Category `gorm:"foreignKey:CategoryId"`
}
