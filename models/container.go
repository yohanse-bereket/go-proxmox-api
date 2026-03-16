package models

import (
	"time"

	"gorm.io/gorm"
)

type Container struct {
	ID        uint           `gorm:"primaryKey"`
	UUID      string         `gorm:"uniqueIndex"`
	Name      string
	CPU       int
	Memory    int
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}