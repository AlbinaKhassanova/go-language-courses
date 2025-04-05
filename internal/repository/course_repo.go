package repository

import (
	"go-language-courses/internal/database"
	"go-language-courses/internal/models"
)

func CreateCourse(course *models.Course) error {
	return database.DB.Create(course).Error
}

func GetCourses() ([]models.Course, error) {
	var courses []models.Course
	err := database.DB.Find(&courses).Error
	return courses, err
}

func GetCourseByID(id uint) (models.Course, error) {
	var course models.Course
	result := database.DB.First(&course, id)
	if result.Error != nil {
		return models.Course{}, result.Error
	}
	return course, nil
}

func SaveCourse(course *models.Course) error {
	return database.DB.Save(course).Error
}

func DeleteCourse(id uint) error {
	if err := database.DB.Delete(&models.Course{}, id).Error; err != nil {
		return err
	}
	return nil
}
