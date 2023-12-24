package usecases

import (
	"fmt"
	"test/domains"
)

type MakeOrderInput struct {
	User     domains.User
	Products []domains.ProductID
}

type MakeOrderOutput struct {
	Order domains.Order
}

func (uc UseCases) MakeOrder(in MakeOrderInput) (MakeOrderOutput, error) {
	if err := validateMakeOrder(in); err != nil {
		return MakeOrderOutput{}, fmt.Errorf("%w: %v", ErrValidation, err)
	}

	orderProducts, err := uc.products.GetList(in.Products)
	if err != nil {
		return MakeOrderOutput{}, fmt.Errorf("%w: %v", ErrInternal, err)
	}

	verifiedOrderProducts, err := uc.products.CheckQuantity(orderProducts)
	userOrder := domains.NewOrder(verifiedOrderProducts)

	return MakeOrderOutput{Order: userOrder}, err
}
