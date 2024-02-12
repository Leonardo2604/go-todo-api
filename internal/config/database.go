package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDatabase() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=123456 dbname=todo port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn),  &gorm.Config{})
	return db, err
}