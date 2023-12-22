package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	text = strings.Replace(text, "\r", "", -1)
	text = strings.Replace(text, "\n", "", -1)	
	age, err := strconv.Atoi(text)
	if err != nil {
		panic(err)
	}
	if age >= 18 {
		fmt.Println("You can vote!")
	} else {
		fmt.Println("You can't vote!")
	}

	bCode := "001"
	switch bCode {
		case "001":
			fmt.Printf("Techno")
		case "010":
			fmt.Println("Gundem")
	}
}
 