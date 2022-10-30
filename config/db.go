package config

import (
	"fmt"
	"log"

	"kanban-kita-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	userName = "diva"
	dbName   = "kanban_kita"
	dbPass   = "root"
	dbPort   = "5432"
	dbHost   = "localhost"
	db       *gorm.DB
	err      error
)

func InitDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, userName, dbPass, dbName, dbPort)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Databases Error", err.Error())
	}
	log.Printf("Databases Connected")
	db.Debug().AutoMigrate(models.User{}, models.Task{}, models.Category{})
}

func GetDB() *gorm.DB {
	return db
}
