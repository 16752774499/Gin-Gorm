package models

import "gorm.io/gorm"

type Student struct {
	Id       int       `json:"id"`
	Number   string    `json:"number"`
	Password string    `json:"password"`
	ClassId  int       `json:"class_id"`
	Name     string    `json:"name"`
	Lesson   []*Lesson `gorm:"many2many:lesson_student"`
}

func (Student) TableName() string {
	return "student"
}

func CustomPreloadScope(db *gorm.DB) *gorm.DB {
	return db.Preload("Student", func(db *gorm.DB) *gorm.DB {
		return DB.Order("id DESC")
	})
}
