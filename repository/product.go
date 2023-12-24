package repository

import (
	"test/domains"
)

type ProductRepository interface {
	GetList(pp []domains.ProductID) ([]domains.Product, error)
	CheckQuantity(pp []domains.Product) ([]domains.Product, error)
}
