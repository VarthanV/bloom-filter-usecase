package controllers

type GoTo string

const (
	GotoHomePage      GoTo = "homepage"
	GotoTwoFactorAuth GoTo = "two_factor_auth"
)

type AuthenticationRequest struct {
	LocationID int `json:"location_id"`
	UserID     int `json:"user_id"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type AuthenticateSuccessResponse struct {
	Goto GoTo `json:"goto"`
}

type DropdownOption struct {
	Option string `json:"option"`
	Value  int    `json:"value"`
}
