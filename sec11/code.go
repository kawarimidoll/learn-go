package main

import (
	"context"
	"fmt"
	"time"

	"github.com/pelletier/go-toml"
	"golang.org/x/sync/semaphore"
)

var s *semaphore.Weighted = semaphore.NewWeighted(2)

func longProcess(ctx context.Context) {
	if err := s.TryAcquire(1); !err {
		fmt.Println("Could not get lock!")
		return
	}
	// if err := s.Acquire(ctx, 1); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	fmt.Println("Wait...")
	time.Sleep(1 * time.Second)
	fmt.Println("Done!")
}

func main() {
	conf, err := toml.LoadFile("sample.toml")
	if err != nil {
		fmt.Println("Error", err.Error())
		return
	}

	fmt.Println(conf.Get("title"))
	fmt.Println(conf.Get("owner"))
	fmt.Println(conf.Get("database.server"))
	fmt.Println(conf.Get("clients.data"))
	fmt.Println()

	ctx := context.TODO()
	go longProcess(ctx)
	go longProcess(ctx)
	go longProcess(ctx)
	go longProcess(ctx)

	time.Sleep(5 * time.Second)
}
