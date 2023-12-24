package infrastructure

import "test/domains"

type OrderRepository struct {
	// Database connection or any other infrastructure-specific dependencies
}

func (r *OrderRepository) Save(order domains.Order) error {
	// Implement the logic to save the user in the database
	// we should hash user password before saving
	return nil
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{}
}
