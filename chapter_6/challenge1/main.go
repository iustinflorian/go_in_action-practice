package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var count int64
	// var mutex sync.Mutex

	for range 1000 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				// mutex.Lock()
				atomic.AddInt64(&count, 1) // ❌ DATA RACE HERE!
				// mutex.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Println("Final count:", count)
}
