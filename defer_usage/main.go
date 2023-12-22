package main

import "fmt"

func main() {
	// defer fmt.Println("Hello")
	// defer fmt.Println("Kemal")
	// fmt.Println("World!")

	x := 10
	y := 20

	defer fmt.Printf("x: %d\n", x)
	defer fmt.Printf("y: %d\n", y)

	x = 100
	y = 200

	fmt.Printf("x: %d\n", x)
	fmt.Printf("y: %d\n", y)
}
