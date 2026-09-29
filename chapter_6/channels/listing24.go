package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberCoroutines = 4
	taskLoad         = 10
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	tasks := make(chan string, taskLoad)

	wg.Add(numberCoroutines)
	for gr := 1; gr <= numberCoroutines; gr++ {
		go worker(tasks, gr)
	}

	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("task : %d", post)
	}

	close(tasks)

	wg.Wait()
}

func worker(tasks chan string, worker int) {
	defer wg.Done()

	for {
		task, ok := <-tasks
		if !ok {
			fmt.Printf("worker: %d : shutting down\n", worker)
			return
		}

		fmt.Printf("worker: %d : started %s\n", worker, task)

		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		fmt.Printf("worker: %d : completed %s\n", worker, task)
	}
}
