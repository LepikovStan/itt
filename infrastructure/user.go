package infrastructure

import (
	"test/domains"
	"test/repository"
)

type UserRepository struct {
	// Database connection or any other infrastructure-specific dependencies
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) Save(user domains.User) error {
	// Implement the logic to save the user in the database
	// we should hash user password before saving
	return nil
}

func (r *UserRepository) FindByID(id int) (domains.User, error) {
	// Implement the logic to fetch the user from the database by ID
	return domains.User{}, repository.ErrNotFound
}
