package infrastructure

import "test/domains"

type ProductRepository struct {
	// Database connection or any other infrastructure-specific dependencies
}

func (p ProductRepository) GetList(pp []domains.ProductID) ([]domains.Product, error) {
	// get products list from Database by their IDs list
	panic("implement me")
}

func (p ProductRepository) CheckQuantity(pp []domains.Product) ([]domains.Product, error) {
	// check product quantity
	verifiedProducts := make([]domains.Product, 0)
	for i := 0; i < len(pp); i++ {
		if pp[i].Quantity() == 0 {
			continue
		}
		verifiedProducts = append(verifiedProducts, pp[i])
	}

	return verifiedProducts, nil
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}
