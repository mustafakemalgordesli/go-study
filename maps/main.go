package main

import "fmt"

func main() {
	// var names map[string]int = make(map[string]int, 0)

	// names["Mustafa"] = 1
	// names["Kemal"] = 2
	// names["Murat"] = 3

	// fmt.Println(names)


	names := map[string]int{
		"Mustafa": 1,
		"Kemal": 2,
	}

	delete(names, "Mustafa")

	fmt.Println(names)
}