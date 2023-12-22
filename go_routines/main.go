package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	wg := sync.WaitGroup{}

	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			fmt.Println("f1: ", i)
			time.Sleep(time.Second)
		}
	}()
	go func() {
		defer wg.Done()
		f2()
	}()

	wg.Wait()

	fmt.Println("End of the main")
}

func f2() {
	for i := 0; i < 10; i++ {
		fmt.Println("f2: ", i)
		time.Sleep(time.Second)
	}
}
