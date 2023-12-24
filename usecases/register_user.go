package usecases

import (
	"errors"
	"fmt"
	"test/domains"
	"test/repository"
)

type RegisterUserInput struct {
	Age       int
	IsMarried bool
	Firstname string
	Lastname  string
	Fullname  string
	Password  string
}

type RegisterUserOutput struct {
	User domains.User
}

func (uc UseCases) RegisterUser(in RegisterUserInput) (RegisterUserOutput, error) {
	if err := validateRegisterUser(in); err != nil {
		return RegisterUserOutput{}, fmt.Errorf("%w: %v", ErrValidation, err)
	}

	user := domains.NewUser().
		SetLastName(in.Lastname).
		SetAge(in.Age).
		SetFirstName(in.Firstname).
		SetIsMarried(in.IsMarried).
		SetPassword(in.Password)

	if err := uc.users.Save(user); err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return RegisterUserOutput{}, ErrNotFound
		}

		return RegisterUserOutput{}, fmt.Errorf("%w: %v", ErrInternal, err)
	}

	return RegisterUserOutput{
		User: user,
	}, nil
}
