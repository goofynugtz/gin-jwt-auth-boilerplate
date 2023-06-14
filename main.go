package main

import (
	"os"

	"github.com/gin-gonic/gin"
	routes "github.com/goofynugtz/gin-jwt-auth-boilerplate/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)

	router.GET("/api/v1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for /api/v1"})
	})

	routes.UserRoutes(router)
	router.GET("/api/v2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for /api/v2"})
	})

	router.Run(":" + port)
}
