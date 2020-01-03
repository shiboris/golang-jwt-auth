package main

import (
	"log"
	"os"

	"./auth"
	"./handler"
	"./middleware"
	"github.com/gin-gonic/gin"
)

func sampleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("before logic")
		c.Next()
		log.Println("after logic")
	}
}

func main() {
	port := os.Getenv("PORT")
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	if port == "" {
		port = "8000"
	}

	sampleGroup := r.Group("/")
	sampleGroup.Use(middleware.SampleMiddleware())
	{
		sampleGroup.GET("/", handler.PrivateHandler)
	}
	r.POST("/tokenAuth", auth.LoginHandler)

	r.Run(":8080")
}
