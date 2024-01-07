package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *controller) ListLocations(ctx *gin.Context) {
	var (
		locations = []DropdownOption{
			{
				Value:  1,
				Option: "New York",
			},
			{
				Value:  2,
				Option: "Madrid",
			},
			{
				Value:  3,
				Option: "Mumbai",
			},
		}
	)
	ctx.JSON(http.StatusOK, locations)
}
