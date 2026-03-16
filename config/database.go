package config

import (
	"fmt"
	"log"
	"os"

	"proxmox-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



var DB *gorm.DB

func ConnectDatabase() {

	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	port := os.Getenv("POSTGRES_PORT")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host,
		user,
		password,
		dbname,
		port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	db.AutoMigrate(&models.Container{})

	DB = db
}