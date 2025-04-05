package services

import (
	"go-language-courses/internal/models"
	"go-language-courses/internal/repository"
)

func CreateCourse(course *models.Course) error {
	return repository.CreateCourse(course)
}

func GetCourses() ([]models.Course, error) {
	return repository.GetCourses()
}

func GetCourseByID(id uint) (models.Course, error) {
	course, err := repository.GetCourseByID(id)
	if err != nil {
		return models.Course{}, err
	}
	return course, nil
}
func UpdateCourse(id uint, course *models.Course) (models.Course, error) {
	// Получаем существующий курс из репозитория
	existingCourse, err := repository.GetCourseByID(id)
	if err != nil {
		return models.Course{}, err // Если курс не найден, возвращаем ошибку
	}

	// Обновляем поля курса
	existingCourse.Name = course.Name
	existingCourse.Language = course.Language
	existingCourse.Description = course.Description
	existingCourse.Duration = course.Duration

	if err := repository.SaveCourse(&existingCourse); err != nil {
		return models.Course{}, err
	}

	return existingCourse, nil
}

func DeleteCourse(id uint) error {
	err := repository.DeleteCourse(id)
	if err != nil {
		return err
	}
	return nil
}
