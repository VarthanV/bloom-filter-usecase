package main

import (
	"net/http"

	"github.com/VarthanV/bloom-filter-usecase/backend/internal/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Setup routes
	ctrl := controllers.New()
	r.POST("/authenticate", ctrl.Authenticate)
	r.POST("/verify-auth", ctrl.VerifyAuthentication)
	r.GET("/locations", ctrl.ListLocations)
	r.GET("/users", ctrl.ListUsers)

	// Run server
	r.Use(cors.Default())
	r.Run()
}
