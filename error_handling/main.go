package main

import (
	"fmt"
	// "os"
)

func main() {
	// fileData, err := os.ReadFile("sample.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(string(fileData))
	// }

	// result, err := divide(10, 0)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(result)
	// }

	productService := ProductService{}

	err := productService.Add(Product{id: 1, name: "", price: 3000})

	if err != nil {
		fmt.Println(err)
	}
}

// func divide(x int, y int) (int, error) {
// 	if y == 0 {
// 		return 0, errors.New("payda 0 olamaz")
// 	}
// 	return x / y, nil
// }

type Product struct {
	id    int
	name  string
	price int
}

type ProductService struct {
}

func (productService *ProductService) Add(product Product) error {
	if len(product.name) == 0 {
		return ValidationError{text: "Ürün ismi boş olamaz", code: "1001"}
	}

	if product.price < 0 {
		return ValidationError{text: "Ürün fiyatı 0'dan büyük olmalıdır", code: "1002"}
	}

	fmt.Println("Ürün eklendi")

	return nil
}

type ValidationError struct {
	text string
	code string
}

func (validationError ValidationError) Error() string {
	return fmt.Sprintf("Hata : %s, Kod: %s", validationError.text, validationError.code)
}
