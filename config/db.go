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
	conn := "host=localhost user=postgres password=admin123 dbname=latihan2 port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database", err.Error())
	} else {
		fmt.Println("DB connected")
		DB = db
	}
	db.AutoMigrate(&entity.Login{}, &entity.Register{})
}
