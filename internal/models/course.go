package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Name        string `json:"name"`
	Language    string `json:"language"`
	Description string `json:"description"`
	Duration    int    `json:"duration"`
}
