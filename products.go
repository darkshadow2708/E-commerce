package main

type Product struct {
	Name        string
	Categories  []string
	Price       int
	Description string
}

func Newproducts(name string, categories []string, price int, description string) *Product {
	return &Product{
		Name:        name,
		Categories:  append([]string{}, categories...),
		Price:       price,
		Description: description,
	}

}
