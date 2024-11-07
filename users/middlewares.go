package users

import (
	"authService/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// IsNotLoggedMiddleware assure that the users doesn't send a JWT before accessing next step in the middleware pipeline
func IsNotLoggedMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader != "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Already logged"})
		c.Abort()
		return
	}

	c.Next()
}

// IsLoggedMiddleware assure that the user has a JWT before accessing next step of the middleware pipeline
func IsLoggedMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization header missing"})
		c.Abort()
		return
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid Authorization header format"})
		c.Abort()
		return
	}

	token := parts[1]

	if err := jwt.VerifyToken(token); err == nil {
		c.Next()
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{
		"message": "Unauthorized",
	})
	c.Abort()
}
