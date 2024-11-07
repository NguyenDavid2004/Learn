package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"management_student/controllers"
)

// SetupRouter initializes the Gin router and defines the routes
func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.POST("/register", func(c *gin.Context) { controllers.Register(c, db) })
	r.POST("/login", func(c *gin.Context) { controllers.Login(c, db) })
	r.POST("/forgot-password", func(c *gin.Context) { controllers.ForgotPassword(c, db) })
	r.POST("/reset-password", func(c *gin.Context) { controllers.ResetPassword(c, db) })
	r.Use(AuthMiddleware())

	studentGroup := r.Group("/students")
	{
		// Routes with TeacherOnly middleware (only accessible by teachers)
		studentGroup.POST("/", TeacherOnly, func(c *gin.Context) { controllers.AddStudent(c, db) })
		studentGroup.DELETE("/:id", TeacherOnly, func(c *gin.Context) { controllers.DeleteStudent(c, db) })
		studentGroup.PUT("/:id", TeacherOnly, func(c *gin.Context) { controllers.UpdateStudent(c, db) })
		studentGroup.PATCH("/:id/grade", TeacherOnly, func(c *gin.Context) { controllers.UpdateGrade(c, db) })

		// Routes accessible by all authenticated users
		studentGroup.GET("/:id", func(c *gin.Context) { controllers.GetStudent(c, db) })
		studentGroup.GET("/search", func(c *gin.Context) { controllers.SearchStudents(c, db) })
	}
	subjectGroup := r.Group("/subjects")
	{
		subjectGroup.POST("/", TeacherOnly, func(c *gin.Context) { controllers.AddSubject(c, db) })
		subjectGroup.POST("/assign", TeacherOnly, func(c *gin.Context) { controllers.AssignSubjectToStudent(c, db) })
	}

	return r
}
