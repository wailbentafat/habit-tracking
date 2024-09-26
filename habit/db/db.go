package db

import (
	"habit/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
func SetupDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("habit-tracker.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = DB.AutoMigrate(&models.User{}, &models.Habit{}, &models.Progres{})
	if err != nil {
		log.Fatal("Failed to auto-migrate database:", err)
	}
	models.SeedCategories(DB)
}
