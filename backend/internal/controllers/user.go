package controllers

import (
	"fmt"
	"net/http"

	"github.com/VarthanV/bloom-filter-usecase/backend/internal/bloomfilter"
	"github.com/gin-gonic/gin"
)

type controller struct {
	filter                bloomfilter.IBloomFilter
	userAuthActivityTrack map[int]map[int]bool
}

func New() *controller {
	return &controller{
		filter:                bloomfilter.New(10),
		userAuthActivityTrack: map[int]map[int]bool{},
	}
}

func (c *controller) Authenticate(ctx *gin.Context) {
	var (
		request = AuthenticationRequest{}
	)

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		fmt.Println("unable to bind request ", err.Error())
		ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	err = request.Validate()
	if err != nil {
		fmt.Println("error in validating request ", err.Error())
		ctx.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

}
