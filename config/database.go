package config

import (
	"log"
	"os"

	"proxmox-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



var DB *gorm.DB

func ConnectDatabase() {

	dsn := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	db.AutoMigrate(&models.Container{})

	DB = db
}