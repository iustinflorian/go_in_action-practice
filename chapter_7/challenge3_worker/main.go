package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	wg.Add(3)
	for range 3 {
		go worker(jobs, results, &wg)
	}

	for i := 1; i <= 10; i++ {
		jobs <- i
	}
	close(jobs)

	wg.Wait()
	close(results)

	for result := range results {
		fmt.Println("result:", result)
	}
}

func worker(jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	var result int

	for job := range jobs {
		result = job * 2
		results <- result
	}
}
