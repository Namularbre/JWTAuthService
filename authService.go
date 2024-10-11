package main

import (
	"authService/users"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	r := gin.Default()

	r.POST("/login", users.Login)
	r.POST("/register", users.PostCreate)
	r.PUT("/update", users.IsLoggedMiddleware, users.PutUser)
	r.DELETE("/delete", users.IsLoggedMiddleware, users.DeleteUser)

	listeningAddress := os.Getenv("ADDRESS") + ":" + os.Getenv("PORT")

	err := r.Run(listeningAddress)
	if err != nil {
		log.Fatalln(err)
	}
}
