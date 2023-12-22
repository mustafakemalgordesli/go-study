package main

import "fmt"

func main() {
	// var productName string
	// var quantity int
	// var discount float32
	// var isInStock bool

	// productName = "Golang Dersleri"
	// quantity = 10
	// discount = 10.5
	// isInStock = true

	// fmt.Println(productName, reflect.TypeOf(productName))
	// fmt.Println(quantity)
	// fmt.Println(discount, reflect.TypeOf(discount))
	// fmt.Println(isInStock, reflect.TypeOf(isInStock))

	// productName := "Golang"
	// fmt.Println(productName)

	//Format
	var productName string
	var quantity int
	var discount float32
	var isInStock bool

	productName = "Golang Dersleri"
	quantity = 1000
	discount = 10.5
	isInStock = true

	fmt.Printf("Product Name: %s, Quantity: %d, Discount: %f, IsInStock: %t", productName, quantity, discount, isInStock)
}