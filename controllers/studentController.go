package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"management_student/models"
	"net/http"
	"strconv"
)

func AddStudent(c *gin.Context, db *gorm.DB) {
	var studentData struct {
		Name    string  `json:"name"`
		Age     int     `json:"age"`
		ClassID uint    `json:"class_id"`
		Grade   float64 `json:"grade"`
	}

	// Bind JSON input to studentData
	if err := c.ShouldBindJSON(&studentData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the specified ClassID exists
	var class models.Class
	if err := db.First(&class, studentData.ClassID).Error; err != nil {
		// If ClassID does not exist, create a new class
		newClass := models.Class{Name: fmt.Sprintf("Class %d", studentData.ClassID)}
		if err := db.Create(&newClass).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create class"})
			return
		}
		class = newClass
	}

	// Proceed to add the student with the ClassID
	student := models.Student{
		Name:    studentData.Name,
		Age:     studentData.Age,
		ClassID: class.ID,
		Grade:   studentData.Grade,
	}

	// Add student to the database
	if err := models.AddStudent(db, &student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add student"})
		return
	}

	c.JSON(http.StatusOK, student)
}

func DeleteStudent(c *gin.Context, db *gorm.DB) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := models.DeleteStudent(db, uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete student"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Student deleted successfully"})
}

func UpdateStudent(c *gin.Context, db *gorm.DB) {
	// Retrieve the student ID from the URL parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	// Bind the JSON body to a Student struct
	var updatedData struct {
		Name    string `json:"name"`
		Age     int    `json:"age"`
		ClassID uint   `json:"class_id"`
		//Grade   float64 `json:"grade"`
	}

	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Find the existing student by ID
	var student models.Student
	if err := db.First(&student, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	// Update the student's fields
	student.Name = updatedData.Name
	student.Age = updatedData.Age
	student.ClassID = updatedData.ClassID
	//student.Grade = updatedData.Grade

	// Save the changes
	if err := db.Save(&student).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update student"})
		return
	}

	c.JSON(http.StatusOK, student)
}

func GetStudent(c *gin.Context, db *gorm.DB) {
	id, _ := strconv.Atoi(c.Param("id"))
	student, err := models.GetStudent(db, uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	c.JSON(http.StatusOK, student)
}

func SearchStudents(c *gin.Context, db *gorm.DB) {
	name := c.Query("name")
	students, err := models.SearchStudents(db, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search students"})
		return
	}
	c.JSON(http.StatusOK, students)
}

func UpdateGrade(c *gin.Context, db *gorm.DB) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	var gradeData struct {
		Grade float64 `json:"grade"`
	}

	if err := c.ShouldBindJSON(&gradeData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Update the student's grade in the database
	if err := models.UpdateStudentGrade(db, uint(id), gradeData.Grade); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update grade"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Grade updated successfully"})
}
