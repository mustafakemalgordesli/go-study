package main

import (
	"fmt"
)

func main() {
	// names := []string { "Mustafa", "Kemal", "Murat"}

	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	// for i := 0; i <= 10; i++ {
	// 	if i == 3 {
	// 		continue
	// 	}
	// 	if i == 5 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	// var numbers = []int{1, 2, 3, 4}

	// for _, value := range numbers {
	// 	fmt.Println(value)
	// }

	// var lang = "Golang"
	// for _, character := range lang {
	// 	fmt.Printf("Character %c\n", character)
	// }

	names := map[string]int{
		"Mustafa": 1,
		"Kemal":   2,
	}
	for key, value := range names {
		fmt.Println(key, " ", value)
	}
}
