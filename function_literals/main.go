package main

import (
	"fmt"
)

func main() {
	// func() {
	// 	fmt.Println("f1")
	// }()

	// add := func(a int, b int) int {
	// 	return a + b
	// }

	// fmt.Println(reflect.TypeOf((add)))

	fmt.Println(calculator(3, 5, func(a int, b int) int { return a - b }))
}

func calculator(x int, y int, f1 func(int, int) int) int {
	return f1(x, y)
}
