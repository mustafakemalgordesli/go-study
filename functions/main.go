package main

import "fmt"

func main() {
	// fmt.Println(add(3, 5))
	// total, difference, multiply := calculation(10, 5)
	// fmt.Println(total, " ", difference, " ", multiply)
	yazdir(sum(3, 4, 5, 5))
}

func sum(numbers ...int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}

// func calculation(x int, y int) (int, int, int) {
// 	return x + y, x - y, x * y
// }

// func add(a int, b int) int {
// 	return a + b
// }

func yazdir(a any) {
	fmt.Println(a)
}
