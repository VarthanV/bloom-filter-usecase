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

func (c *controller) getKey(userID int, locationID int) string {
	return fmt.Sprintf("user_%d:location_%d", userID, locationID)
}

func (c *controller) ListUsers(ctx *gin.Context) {
	var (
		users = []DropdownOption{
			{
				Value:  1,
				Option: "John",
			},
			{
				Value:  2,
				Option: "Simon",
			},
			{
				Value:  3,
				Option: "Carla",
			},
		}
	)
	ctx.JSON(http.StatusOK, users)

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

	// Check if the key exists in filter
	// user_id:location_id

	key := c.getKey(request.UserID, request.LocationID)

	isExists, err := c.filter.CheckMembership(key)
	if err != nil {
		fmt.Println("error in checking if key exists ", err)
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{
			Message: "Internal Server error",
		})
		return
	}

	// Might be a chance for false positive
	if isExists {
		recordMap, ok := c.userAuthActivityTrack[request.UserID]
		if !ok {
			fmt.Println("record for user doesnt exist in map")
			ctx.JSON(http.StatusOK, AuthenticateSuccessResponse{
				Goto: GotoTwoFactorAuth,
			})
			return
		}

		hasUserAuthenticatedInLocation, ok := recordMap[request.LocationID]
		if !ok {
			fmt.Println("record for user location doesnt exist in map")
			ctx.JSON(http.StatusOK, AuthenticateSuccessResponse{
				Goto: GotoTwoFactorAuth,
			})
			return
		}

		if hasUserAuthenticatedInLocation {
			ctx.JSON(http.StatusOK, AuthenticateSuccessResponse{
				Goto: GotoHomePage,
			})
			return
		}
		ctx.JSON(http.StatusOK, AuthenticateSuccessResponse{
			Goto: GotoTwoFactorAuth,
		})
		return

	}

	ctx.JSON(http.StatusOK, AuthenticateSuccessResponse{
		Goto: GotoTwoFactorAuth,
	})

}

func (c *controller) VerifyAuthentication(ctx *gin.Context) {

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

	_, ok := c.userAuthActivityTrack[request.UserID]
	if !ok {
		c.userAuthActivityTrack[request.UserID] = map[int]bool{}
	}

	c.userAuthActivityTrack[request.UserID][request.LocationID] = true

	key := c.getKey(request.UserID, request.LocationID)
	err = c.filter.Add(key)
	if err != nil {
		fmt.Println("unable to add key in filter ", err)
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{
			Message: "Internal Server error",
		})
		return
	}

	ctx.JSON(http.StatusOK, AuthenticateSuccessResponse{
		Goto: GotoHomePage,
	})

}
