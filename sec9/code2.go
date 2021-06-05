package main

import (
	"context"
	"fmt"
	"time"
)

const (
	c1 = iota
	c2
	c3
)

const (
	_      = iota
	KB int = 1 << (10 * iota)
	MB
	GB
)

func longProcess(ctx context.Context, ch chan string) {
	fmt.Println("longProcess: start")
	time.Sleep(2 * time.Second)
	fmt.Println("longProcess: end")
	ch <- "done"
}

func main() {
	fmt.Println(c1, c2, c3)
	fmt.Println(KB, MB, GB)

	ch := make(chan string)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	go longProcess(ctx, ch)

	for {
		select {
		case <-ch:
			fmt.Println("success")
			return
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		}
	}
}
