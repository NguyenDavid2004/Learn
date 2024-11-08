package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"management_student1/models"
	"net/http"
)

func AddSubject(c *gin.Context, db *gorm.DB) {
	var subjectData struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	if err := c.ShouldBind(&subjectData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	subject := models.Subject{
		Name:        subjectData.Name,
		Description: subjectData.Description,
	}
	if err := models.CreateSubject(db, &subject); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create subject"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": subject})
}

func AssginSubjectToStudent(c *gin.Context, db *gorm.DB) {
	var data struct {
		StudentID uint `json:"student_id"`
		SubjectID uint `json:"subject_id"`
	}
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	studentSubject := models.StudentSubject{
		StudentID: data.StudentID,
		SubjectID: data.SubjectID,
	}

	if err := db.Create(&studentSubject).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to assgin subject"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": studentSubject})
}
