package models

import "gorm.io/gorm"

type Subject struct {
	ID          uint   `gorm:"primary_key"`
	Name        string `gorm:"not null"`
	Description string
}

type StudentSubject struct {
	StudentID uint `gorm:"primary_key"`
	SubjectID uint `gorm:"primary_key"`
}

func CreateSubject(db *gorm.DB, subject *Subject) error {
	return db.Create(subject).Error
}
