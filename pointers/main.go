package main

import "fmt"

func main() {
	// var a int
	// var p *int

	// a = 10
	// p = &a

	// *p = 20

	// fmt.Println(a)
	// fmt.Println(p)
	// fmt.Println(*p)

	// var a = 10
	// var b int
	// var p *int
	// b = a
	// p = &a
	// *p = 20
	// fmt.Println(a, b)

	var a int = 10
	add12(&a)
	fmt.Println(a)
}

func add12(x *int) {
	*x = *x + 12
}
