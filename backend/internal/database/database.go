package database

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

func NewDatabase(dsn string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
} 