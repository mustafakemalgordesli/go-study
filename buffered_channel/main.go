package main

import (
	"fmt"
	"time"
)

func main() {
	// myChannel := make(chan int, 2)
	// myChannel <- 1
	// myChannel <- 2
	// fmt.Println(<-myChannel)
	// fmt.Println(<-myChannel)
	// myChannel <- 3
	// fmt.Println(<-myChannel)

	// messages := make(chan string, 2)

	// go func() {
	// 	data1 := <-messages
	// 	fmt.Println("First : ", data1)
	// 	data2 := <-messages
	// 	fmt.Println("Second : ", data2)
	// }()

	// messages <- "Hello"
	// messages <- "World"
	// messages <- "World2"
	// time.Sleep(1 * time.Second)

	bufferedChannel := make(chan int, 4)

	go func() {
		for i := 1; i <= 10; i++ {
			bufferedChannel <- i
			fmt.Println("Send data:", i)
		}
		close(bufferedChannel)
	}()

	time.Sleep(3 * time.Second)

	for data := range bufferedChannel {
		fmt.Println("Received data: ", data)
		time.Sleep(1 * time.Second)
	}

	fmt.Println("End of the main")
}
