package main

import (
	"authService/migration"
	"authService/users"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	migration.Init()

	r := gin.Default()

	err := r.SetTrustedProxies([]string{})
	if err != nil {
		panic(err)
	}

	r.Use(func(c *gin.Context) {
		start := time.Now()
		c.Next()

		latency := time.Since(start)
		log.Printf("Request processed in %s", latency)

		status := c.Writer.Status()
		log.Printf("Status code: %d", status)
	})

	r.POST("/login", users.IsNotLoggedMiddleware, users.Login)
	r.POST("/register", users.IsNotLoggedMiddleware, users.Register)
	r.POST("/authenticate", users.IsLoggedMiddleware, users.Authenticate)
	r.GET("/isAdmin", users.IsLoggedMiddleware, users.IsAdmin)

	listeningAddress := os.Getenv("ADDRESS") + ":" + os.Getenv("PORT")

	log.Printf("Listening on http://%s", listeningAddress)

	err = r.Run(listeningAddress)
	if err != nil {
		log.Fatalln(err)
	}
}
