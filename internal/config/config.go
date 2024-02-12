package config

import "gorm.io/gorm"

var (
	db *gorm.DB
)

func InitializeConfig() {
	var err error
	
	db, err = InitializeDatabase();

	if err != nil {
		panic("Failed on database connection")
	}
}

func GetDB() *gorm.DB {
	return db
}