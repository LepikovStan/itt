package usecases

import "errors"

var (
	ErrValidation = errors.New("validation error")
	ErrNotFound   = errors.New("not found")
	ErrInternal   = errors.New("internal error")
)
