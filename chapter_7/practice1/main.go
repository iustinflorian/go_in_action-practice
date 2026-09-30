package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	data, err := fetchData(ctx)
	if err == nil {
		fmt.Println(data)
	} else {
		fmt.Println(err)
	}
}

func fetchData(ctx context.Context) (string, error) {
	ch := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "data loaded"
	}()

	select {
	case res := <-ch:
		return res, nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}
