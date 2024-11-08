package models

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm: "primary_key"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Student struct {
	ID      uint    `gorm: "primary_key"`
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	ClassId uint    `json:"class_id"`
	Grade   float64 `json:"grade"`
}

type Class struct {
	ID      uint      `gorm: "primary_key"`
	Name    string    `json:"name"`
	Student []Student `json:"student"`
}

func AddStudent(db *gorm.DB, student *Student) error {
	return db.Create(student).Error
}

func DeleteStudent(db *gorm.DB, id uint) error {
	return db.Delete(&Student{}, id).Error
}

func GetStudent(db *gorm.DB, id uint) (*Student, error) {
	var student Student
	err := db.First(&student, id).Error
	return &student, err
}

func SearchStudents(db *gorm.DB, name string) ([]Student, error) {
	var students []Student
	err := db.Where("name Like ?", "%"+name+"%").Find(&students).Error
	return students, err
}
func UpdateStudentGrade(db *gorm.DB, id uint, grade float64) error {
	return db.Model(&Student{}).Where("id = ?", id).Update("grade", grade).Error
}

func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func RegisterUser(db *gorm.DB, user *User) error {
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return db.Create(user).Error
}

func AuthenticateUser(db *gorm.DB, email string, password string) (*User, error) {
	var user User
	err := db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, errors.New("invalid email or password")
	}
	if !CheckPasswordHash(password, user.Password) {
		return nil, errors.New("invalid password")
	}
	return &user, nil
}
