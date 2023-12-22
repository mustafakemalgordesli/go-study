package main

import "fmt"

func main() {
	customer1 := Customer{id: 1, name: "Mustafa Kemal Gordesli", age: 30}
	customer1.changeName("Mustafa")
	customer1.ToString()
}

type Customer struct {
	id   int32
	name string
	age  int
}

func (customer *Customer) ToString() {
	fmt.Printf("Id: %d, Name: %s, Age: %d\n", customer.id, customer.name, customer.age)
}

func (customer *Customer) changeName(newName string) {
	customer.name = newName
}
