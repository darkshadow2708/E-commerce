package main

type OrderItems struct {
	Prdct    []*Product
	Quantity int
	Tax      float64
	Total    float64
}

func NewOrderItem(prdct []*Product, quantity int, t float64, total float64) *OrderItems {
	return &OrderItems{
		Prdct:    append([]*Product{}, prdct...),
		Quantity: quantity,
		Tax:      t,
		Total:    total,
	}
}
