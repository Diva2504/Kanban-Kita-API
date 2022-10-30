package controller

import "gorm.io/gorm"

type Handlers struct {
  Connect *gorm.DB
}
