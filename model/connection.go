package model

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConnection() *gorm.DB {
	dsn := os.Getenv("IPO_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Ipo{})
	return db

}
