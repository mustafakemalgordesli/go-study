package main

import "fmt"

func main() {

	//Fixed
	// var names [4]string

	// fmt.Println(names)

	// names = [4]string{ "Merhaba", "Dünyalı", "Kemal", "Deneme"}
	// fmt.Println(names[1:2])

	var names = []string{"Mustafa", "Kemal"}
	fmt.Println(names)
	names = append(names, "Gördesli")
	fmt.Println(names)
}