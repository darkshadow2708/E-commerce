package main

type Order struct {
	user            User
	oi              []*OrderItems
	ShippingAddress string
}

func NewOrder(u User, o []*OrderItems, address string) *Order {
	return &Order{
		user:            u,
		oi:              append([]*OrderItems{}, o...),
		ShippingAddress: address,
	}
}
func (o *Order) TotalwithoutTax() int {
	var m int = 0
	for i := 0; i < len(o.oi); i++ {
		m += o.oi[i].Prdct[i].Price * o.oi[i].Quantity
	}
	return m

}

func (o *Order) TotalwithTax() float64 {
	var m int = 0
	for i := 0; i < len(o.oi); i++ {
		m += o.oi[i].Prdct[i].Price * o.oi[i].Quantity
	}
	var n float64 = 0
	n = float64(m) + float64(m)*(o.oi.Tax)/100
	o.oi.Total = n
	return n

}
