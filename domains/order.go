package domains

import (
	"github.com/google/uuid"
)

type Order struct {
	id       uuid.UUID
	products []Product
}

func (o Order) ID() uuid.UUID {
	return o.id
}

func NewOrder(pp []Product) Order {
	return Order{
		id:       uuid.New(),
		products: pp,
	}
}
