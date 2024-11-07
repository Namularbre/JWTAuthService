package main

import (
	"authService/migration"
	"authService/users"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
	"time"
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

	listeningAddress := os.Getenv("ADDRESS") + ":" + os.Getenv("PORT")

	log.Printf("Listening on http://" + listeningAddress)

	err = r.Run(listeningAddress)
	if err != nil {
		log.Fatalln(err)
	}
}
