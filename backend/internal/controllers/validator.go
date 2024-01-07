package controllers

import "errors"

var (
	ErrInvalidLocationID = errors.New("invalid location_id")
	ErrInvalidUserID     = errors.New("invalid user_id")
)

func (r *AuthenticationRequest) Validate() error {
	if r.LocationID <= 0 {
		return ErrInvalidLocationID
	}
	if r.UserID <= 0 {
		return ErrInvalidUserID
	}
	return nil
}
