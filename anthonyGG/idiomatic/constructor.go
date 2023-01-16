package main

type Order struct {
	Size float64
}

func NewOrder(size float64) *Order {
	return &Order{
		Size: size,
	}
}
