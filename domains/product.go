package domains

type ProductID int

type Product struct {
	id          ProductID
	description string
	tags        []string
	quantity    int
}

func (p Product) Quantity() int {
	return p.quantity
}
