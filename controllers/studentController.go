package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"management_student1/models"
	"net/http"
	"strconv"
)

func AddStudent(c *gin.Context, db *gorm.DB) {
	var studentData struct {
		Name    string  `json:"name"`
		Age     int     `json:"age"`
		ClassId int     `json:"class_id"`
		Grade   float64 `json:"grade"`
	}
	if err := c.ShouldBind(&studentData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var class models.Class
	if err := db.First(&class, studentData.ClassId).Error; err != nil {
		newClass := models.Class{Name: fmt.Sprintf("%d", studentData.ClassId)}
		if err := db.Create(&newClass).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create class"})
		}
		class = newClass
	}

	student := models.Student{
		Name:    studentData.Name,
		Age:     studentData.Age,
		ClassId: class.ID,
		Grade:   studentData.Grade,
	}
	if err := models.AddStudent(db, &student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add student"})
		return
	}
	c.JSON(http.StatusOK, student)
}

func DeleteStudent(c *gin.Context, db *gorm.DB) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	if err := models.DeleteStudent(db, uint(id)); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete student"})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Operation Success"})
}

func UpdateStudent(c *gin.Context, db *gorm.DB) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}
	var updateData struct {
		Name    string `json:"name"`
		Age     int    `json:"age"`
		ClassID uint   `json:"class_id"`
	}

	if err := c.ShouldBind(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var student models.Student
	if err := db.First(&student, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Student not found"})
		return
	}

	student.Name = updateData.Name
	student.Age = updateData.Age
	student.ClassId = updateData.ClassID
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Student not found"})
		return
	}
	c.JSON(http.StatusOK, student)
}

func SearchStudent(c *gin.Context, db *gorm.DB) {
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

	if err := c.ShouldBind(&gradeData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := models.UpdateStudentGrade(db, uint(id), gradeData.Grade); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update grade"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Grade updated successfully"})
}
