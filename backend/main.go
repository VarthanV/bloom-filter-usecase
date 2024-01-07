package main

import (
	"net/http"

	"github.com/VarthanV/bloom-filter-usecase/backend/internal/controllers"
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

	// Run server
	r.Run()
}
