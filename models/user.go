package models

import (
	"github.com/asaskevich/govalidator"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
  gorm.Model
  FullName string `json:"full_name"`
  Email string
  Password string
  Role string
  Task []Task
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
  salt := 10
  _, err := govalidator.ValidateStruct(u)
  if err != nil {
    return err
  }
  password := []byte(u.Password)
  hash, _ := bcrypt.GenerateFromPassword(password, salt)
  u.Password = string(hash)
  return nil
}
