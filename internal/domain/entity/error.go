package entity

import "errors"

var (
	ErrInvalidToken = errors.New("invalid or missing API key")
)
