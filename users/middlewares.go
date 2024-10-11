package users

import (
	"authService/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

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
