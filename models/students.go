package models

import (
	"errors"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

func RegisterUser(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}
func UpdateStudentGrade(db *gorm.DB, id uint, grade float64) error {
	return db.Model(&Student{}).Where("id = ?", id).Update("grade", grade).Error
}

func AuthenticateUser(db *gorm.DB, email string, password string) (*User, error) {
	var user User
	err := db.Where("email = ? and password = ? ", email, password).First(&user).Error
	if err != nil {
		return nil, errors.New("invalid email or password")
	}
	return &user, nil
}

type Student struct {
	ID      uint    `gorm:"primaryKey"`
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	ClassID uint    `json:"class_id"`
	Grade   float64 `json:"grade"`
}

type Class struct {
	ID       uint      `gorm:"primaryKey"`
	Name     string    `json:"name"`
	Students []Student `gorm:"foreignKey:ClassID"`
}

func AddStudent(db *gorm.DB, student *Student) error {
	return db.Create(student).Error
}

func DeleteStudent(db *gorm.DB, id uint) error {
	return db.Delete(&Student{}, id).Error
}

func UpdateStudent(db *gorm.DB, student *Student) error {
	return db.Save(student).Error
}

func GetStudent(db *gorm.DB, id uint) (*Student, error) {
	var student Student
	err := db.First(&student, id).Error
	return &student, err
}

func SearchStudents(db *gorm.DB, name string) ([]Student, error) {
	var students []Student
	err := db.Where("name LIKE ?", "%"+name+"%").Find(&students).Error
	return students, err
}
