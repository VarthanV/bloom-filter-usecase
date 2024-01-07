package controllers

type AuthenticationRequest struct {
	LocationID int `json:"location_id"`
	UserID     int `json:"user_id"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

