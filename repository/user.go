package repository

import (
	"errors"
	"test/domains"
)

var (
	ErrNotFound = errors.New("not found")
)

type UserRepository interface {
	Save(user domains.User) error
	FindByID(id int) (domains.User, error)
}
