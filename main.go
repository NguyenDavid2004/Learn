package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"management_student/models"
	"management_student/routes"
)

func main() {
	// Connection string for PostgreSQL
	dsn := "host=localhost user=myuser password=mypassword dbname=student_db port=5434 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db.AutoMigrate(&models.Student{}, &models.Class{})

	r := routes.SetupRouter(db)

	if err := r.Run(":8080"); err != nil {
		fmt.Println("Failed to run server:", err)
	}
}
