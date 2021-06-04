package main

import (
	"fmt"
	"sync"
)

func producer(c chan int, i int) {
	c <- i * 2
}

func consumer(c chan int, wg *sync.WaitGroup) {
	for i := range c {
		func() {
			defer wg.Done()
			fmt.Println("process", i*1000)
		}()
	}
}

func main() {
	var wg sync.WaitGroup

	c := make(chan int)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go producer(c, i)
	}

	go consumer(c, &wg)
	wg.Wait()
	close(c)
	fmt.Println("done")
}
