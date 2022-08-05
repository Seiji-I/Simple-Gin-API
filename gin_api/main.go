package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"time"
)

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"http://localhost:3000"},
    AllowMethods:     []string{"GET", "POST", "OPTIONS"},
    AllowHeaders:     []string{"Origin"},
    ExposeHeaders:    []string{"Content-Length"},
    AllowCredentials: true,
    MaxAge: 1 * time.Hour,
  }))

	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H {
			"message": "GET METHOD: OK",
		})
	})
	router.POST("/test", func(c *gin.Context) {
		c.JSON(200, gin.H {
			"message": "OK",
		})
	})

	router.Run(":3333")
}