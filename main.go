package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
		fmt.Println(port)
	}

	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/api1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api1"})
	})

	router.GET("/api2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api2"})
	})

	router.Run(":" + port)
}
