package main

import "fmt"

func main() {
	x := NewInventory()
	s1 := []string{"Shoes", "Footwears"}
	p1 := Newproducts("Reebok", s1, 5000, "Light Weight")
	p2 := Newproducts("Adidas", s1, 5000, "Light Weight")
	x.AddProduct(p1)
	x.AddProduct(p1)
	x.AddProduct(p2)
	fmt.Println(x.Quantity(p1))
	x.RemoveProduct(p1)
	fmt.Println(x.Quantity(p1))

}
