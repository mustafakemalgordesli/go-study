package main

import (
	"context"
	"fmt"
	"time"
)

type Product struct {
	Id   int64
	Name string
}

var productChannel = make(chan Product)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	go getProductDetailsFromExternalService(10)

	select {
	case productDetail := <-productChannel:
		fmt.Println("Product Details: ", productDetail)
	case <-ctx.Done():
		fmt.Println("Timeout occured, context cancelled")
	}

	fmt.Println("End of the main")
}

func getProductDetailsFromExternalService(productId int64) {
	time.Sleep(5 * time.Second)
	productChannel <- Product{
		Id:   productId,
		Name: "Laptop",
	}
}
