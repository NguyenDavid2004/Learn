package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"management_student/controllers"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
			c.Abort()
			return
		}
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		claims := &controllers.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return controllers.JwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Set user context
		c.Set("user", claims.Email)
		c.Set("role", claims.Role)
		c.Next()
	}
}

func TeacherOnly(c *gin.Context) {
	role, exists := c.Get("role")
	if !exists || role != "teacher" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied: Teacher only"})
		c.Abort()
		return
	}
	c.Next()
}
