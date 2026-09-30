package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan string)

	wg.Add(1)
	go producer(ch, &wg)

	for {
		val, ok := <-ch
		if ok != true {
			fmt.Println("channel closed")
			break
		}
		fmt.Println("Received:", val)
	}

	wg.Wait()
}

func producer(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- "golang"
	ch <- "channels"
	ch <- "concurrency"
	close(ch)
}
