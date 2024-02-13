package db

import (
	"github.com/Leonardo2604/go-todo-api/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() error {
	c := config.GetConfig()

	var err error

	db, err = gorm.Open(postgres.Open(c.GetDatabase().GetDSN()), &gorm.Config{})

	return err
}

func GetDB() *gorm.DB {
	return db
}
