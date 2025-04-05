package database

import (
	"go-language-courses/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() *gorm.DB {

	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}

	// Автоматическое создание таблиц на основе моделей
	DB.AutoMigrate(&models.Course{}, &models.Student{}, &models.Enrollment{})
	return DB
}
