package client

import "errors"

var (
	ErrInvalidHotel         = errors.New("invalid hotel")
	ErrInvalidUniqueID      = errors.New("invalid unique ID")
	ErrUnexpectedStatusCode = errors.New("unexpected status code")
	ErrHabboNotFound        = errors.New("habbo not found")
)
