package main

import (
	"fmt"
	"sync"
)

func main(){
	var wg sync.WaitGroup
	ch := make(chan string)

	wg.Add(1)
	producer(ch, &wg)

	for {
		val, ok := <-ch
		if ok == true{
			fmt.Println("channel closed")
			return
		}
		fmt.Println("Received:", val)
	}
}

func producer(ch chan string, wg *sync.WaitGroup){
	defer wg.Wait()
	ch <- "golang"
	ch <- "channels"
	ch <- "concurrency"
	close(ch)
}