package usecases

import (
	"errors"
	"test/product"
	"unicode/utf8"
)

func validateRegisterUser(input RegisterUserInput) error {
	if err := validateUserAge(input.Age); err != nil {
		return err
	}

	if err := validateUserPassword(input.Password); err != nil {
		return err
	}

	return nil
}

func validateMakeOrder(input MakeOrderInput) error {
	if err := validateOrderProductsLen(input.Products); err != nil {
		return err
	}

	return nil
}

func validateOrderProductsLen(pp []product.ID) error {
	if len(pp) == 0 {
		return errors.New("products list must not be empty")
	}

	return nil
}

func validateUserAge(age int) error {
	if age < 18 {
		return errors.New("user age less then 18")
	}

	return nil
}

func validateUserPassword(password string) error {
	if utf8.RuneCountInString(password) < 8 {
		return errors.New("user password less then 8")
	}

	return nil
}
