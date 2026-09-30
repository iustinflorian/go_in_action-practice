package main

import (
	"fmt"
	"sync"
)

func main() {
	channel := make(chan int, 5)
	var wg sync.WaitGroup

	wg.Add(2)

	go producer(channel, &wg)
	go consumer(channel, &wg)

	wg.Wait()
	fmt.Println("done")
}

func producer(channel chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 5; i++ {
		channel <- i
	}

	close(channel)
}

func consumer(channel chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for value := range channel {
		fmt.Println("Received:", value)
	}
}
