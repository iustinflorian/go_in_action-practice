package main

import (
	"fmt"
	"sync"
)

func main() {
	var producerwg sync.WaitGroup
	var consumerwg sync.WaitGroup
	jobs := make(chan int)

	producerwg.Add(2)
	go producer1(jobs, &producerwg)
	go producer2(jobs, &producerwg)

	consumerwg.Add(1)
	go consumer(jobs, &consumerwg)

	go func(){
		producerwg.Wait()
		close(jobs)
	}()

	consumerwg.Wait()
	fmt.Println("done")
}

func producer1(jobs chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 10; i <= 14; i++ {
		jobs <- i
	}
}

func producer2(jobs chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 20; i <= 24; i++ {
		jobs <- i
	}
}

func consumer(jobs chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range jobs {
		fmt.Println("Received:", val)
	}
}
