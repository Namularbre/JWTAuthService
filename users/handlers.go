package users

import (
	"authService/hashing"
	"authService/jwt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Register(c *gin.Context) {
	user := User{}

	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}

	if existingUser, err := SelectByUsername(user.Username); err == nil && existingUser == nil {
		user, err := Create(user.Username, user.Password)
		if err != nil {
			log.Printf("Error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})
			return
		}
		log.Printf("Created user %v", user)
		c.JSON(http.StatusOK, gin.H{
			"user": user,
		})
	} else if err == nil && existingUser != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "User already exists",
		})
	} else {
		log.Printf("Error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
	}
}

func Login(c *gin.Context) {
	var user *User

	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}

	requestPassword := user.Password
	user, err := SelectByUsername(user.Username)
	if err != nil {
		log.Printf("Error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	if hashing.Compare(requestPassword, user.Password) {
		token, err := jwt.CreateToken(user.Username, user.IdUser)
		if err != nil {
			log.Printf("Error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
	}
}

func Authenticate(c *gin.Context) {
	var token string

	if err := c.Bind(&token); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}

	if err := jwt.VerifyToken(token); err != nil {

		log.Printf("Error: %v", err)
		c.JSON(http.StatusForbidden, gin.H{
			"logged": false,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"logged": true,
		})
	}
}

func PutUser(c *gin.Context) {
	var username string
	var password string
	var idUser int

	if err, err1, err2 := c.Bind(&username), c.Bind(&password), c.Bind(&idUser); err != nil && err1 != nil && err2 != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}

	user := &User{
		IdUser:   idUser,
		Username: username,
		Password: string(hashing.Hash(password)),
	}

	err := Update(user)
	if err != nil {
		log.Printf("Error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func DeleteUser(c *gin.Context) {
	var idUser int

	if err := c.Bind(&idUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}

	err := Delete(idUser)
	if err != nil {
		log.Printf("Error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted",
	})
}
