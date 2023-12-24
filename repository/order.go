package repository

import "test/domains"

type OrderRepository interface {
	Save(order domains.Order) error
}
