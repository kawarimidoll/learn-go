package main

import (
	"fmt"
)

func producer(c chan int) {
	defer close(c)
	for i := 0; i < 10; i++ {
		c <- i
	}
}

func multi(c_from <-chan int, c_to chan<- int, num int) {
	defer close(c_to)
	for i := range c_from {
		c_to <- i * num
	}
}

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)

	go producer(c1)
	go multi(c1, c2, 2)
	go multi(c2, c3, 4)
	for result := range c3 {
		fmt.Println(result)
	}
}
