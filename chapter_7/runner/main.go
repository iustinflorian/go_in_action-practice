package main

import (
	"log"
	"os"
	"time"
)

const timeout = 3 * time.Second

func main() {
	log.Println("starting work")

	r := New(timeout)

	r.Add(createTask(), createTask(), createTask())

	if err := r.Start(); err != nil {
		switch err {
		case ErrTimeout:
			log.Println("terminating - timeout")
			os.Exit(1)
		case ErrInterrupt:
			log.Println("terminating - interrupt")
			os.Exit(2)
		}
	}

	log.Println("process end")
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("processor - task #%d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
