package main

type User struct {
	Name    string
	Age     int
	Address string
}

func Newuser(name string, age int, address string) *User {
	return &User{
		Name:    name,
		Age:     age,
		Address: address,
	}
}
