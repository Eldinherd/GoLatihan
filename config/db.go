package config

import (
	"fmt"
	"go-trial-class/entity"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnect() {
	conn := "host=localhost user=postgres password=admin123 dbname=trial_class_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect databasr", err.Error())
	} else {
		fmt.Println("DB connected")
		DB = db
	}
	db.AutoMigrate(&entity.Product{}, &entity.Order{})
}