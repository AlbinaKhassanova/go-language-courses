package models

import "gorm.io/gorm"

type Enrollment struct {
	gorm.Model
	CourseID  uint `json:"course_id"`
	StudentID uint `json:"student_id"`
}
