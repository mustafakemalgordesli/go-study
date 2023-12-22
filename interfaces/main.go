package main

import "fmt"

type IShippable interface {
	ShippingCost() int
}

type Book struct {
	desi int
}

type Electronic struct {
	desi int
}

func (book *Book) ShippingCost() int {
	return 5 + book.desi*2
}

func (electronic *Electronic) ShippingCost() int {
	return 10 + electronic.desi*3
}

func main() {
	// book1 := &Book{desi: 10}
	// book2 := &Book{desi: 20}
	// fmt.Println(book1.ShippingCost())
	// fmt.Println(book2.ShippingCost())
	// books := []Book{*book1, *book2}
	// fmt.Println(calculteTotalShippingCost(books))

	// electronics := []Electronic{{desi: 10}, {desi: 30}}
	// calculteTotalShippingCost(electronics)

	var product1 IShippable = &Book{desi: 10}
	var product2 IShippable = &Electronic{desi: 10}

	products := []IShippable{product1, product2}
	fmt.Println(calculteTotalShippingCost(products))
}

func calculteTotalShippingCost(products []IShippable) int {
	total := 0
	for _, product := range products {
		total += product.ShippingCost()
	}
	return total
}

// func calculteTotalShippingCost(books []Book) int {
// 	total := 0
// 	for _, book := range books {
// 		total += book.ShippingCost()
// 	}
// 	return total
// }
