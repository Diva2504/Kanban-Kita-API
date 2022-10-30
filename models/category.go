package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Type string `gorm:"not null" validate:"required, type"`
	Task []Task
}
