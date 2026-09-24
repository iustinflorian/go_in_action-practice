package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	runtime.GOMAXPROCS(runtime.NumCPU())

	wg.Add(2)

	fmt.Println("creating goroutines")

	go printPrime("A")
	go printPrime("B")

	fmt.Println("waiting")
	wg.Wait()

	fmt.Println("terminating")
}

func printPrime(prefix string) {
	defer wg.Done()

	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}

	fmt.Println("completed", prefix)
}
