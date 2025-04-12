package repository

import (
	"go-language-courses/internal/db"
	"go-language-courses/internal/models"
)

func CreateCourse(course *models.Course) error {
	return db.DB.Create(course).Error
}

func GetCourses() ([]models.Course, error) {
	var courses []models.Course
	err := db.DB.Find(&courses).Error
	return courses, err
}

func GetCourseByID(id uint) (models.Course, error) {
	var course models.Course
	result := db.DB.First(&course, id)
	if result.Error != nil {
		return models.Course{}, result.Error
	}
	return course, nil
}

func SaveCourse(course *models.Course) error {
	return db.DB.Save(course).Error
}

func DeleteCourse(id uint) error {
	if err := db.DB.Delete(&models.Course{}, id).Error; err != nil {
		return err
	}
	return nil
}
