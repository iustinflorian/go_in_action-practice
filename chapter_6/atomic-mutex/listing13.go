package main

import (
	"fmt"
)

var (
// counter int64
// wg sync.WaitGroup
)

func main() {
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()

	fmt.Println("final counter:", counter)
}

//func incCounter(id int) {
//	defer wg.Done()
//
//	for range 2 {
//		atomic.AddInt64(&counter, 1)
//
//		runtime.Gosched()
//	}
//}
