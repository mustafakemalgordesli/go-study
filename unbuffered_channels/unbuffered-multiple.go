package main

import (
	"fmt"
	"sync"
)

func main() {
	// myChannel := make(chan int)

	// go func() {
	// 	for i := 0; i <= 10; i++ {
	// 		myChannel <- i
	// 		fmt.Println("Send data: ", i)
	// 		time.Sleep(1 * time.Second)
	// 	}
	// 	close(myChannel)
	// }()

	// for {
	// 	data, isOpen := <-myChannel
	// 	if !isOpen {
	// 		break
	// 	}
	// 	fmt.Println("Received data: ", data)
	// }

	myChannel := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		data := <-myChannel
		fmt.Println("First Go rountine took data: ", data)
	}()
	go func() {
		defer wg.Done()
		data := <-myChannel
		fmt.Println("Second Go routine took data: ", data)
	}()
	go func() {
		myChannel <- 10
		close(myChannel)
	}()
	wg.Wait()
	fmt.Println("End of the main")
}
