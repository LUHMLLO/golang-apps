package main

import "fmt"

func main() {

	var investment int
	var unitPrice int
	var unitsInStock int

	fmt.Print("How much did you invest?: ")
	fmt.Scanln(&investment)

	fmt.Print("Which would be the price per unit?: ")
	fmt.Scanln(&unitPrice)

	fmt.Print("How many units do you have in stock?: ")
	fmt.Scanln(&unitsInStock)

	var income = unitPrice * unitsInStock
	var margin = income - investment

	fmt.Println("Income amount: ", income)
	fmt.Println("Income margin: ", margin)

	if margin > 0 {
		fmt.Println("Investment it's profitable")
	} else {
		fmt.Println("No Profit")
	}
}
