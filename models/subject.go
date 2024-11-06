package models

import "gorm.io/gorm"

type Subject struct {
	ID          uint   `gorm:"primary_key"`
	Name        string `gorm:"not null"`
	Description string
}

func CreateSubject(db *gorm.DB, subject *Subject) error {
	return db.Create(subject).Error
}

type StudentSubject struct {
	StudentID uint `gorm:"primaryKey"`
	SubjectID uint `gorm:"primaryKey"`
}
