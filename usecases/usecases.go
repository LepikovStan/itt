package usecases

import "test/repository"

type UseCases struct {
	users    repository.UserRepository
	products repository.ProductRepository
	orders   repository.OrderRepository
}

func New(
	users repository.UserRepository,
	products repository.ProductRepository,
	orders repository.OrderRepository,
) UseCases {
	return UseCases{
		users:    users,
		products: products,
		orders:   orders,
	}
}
