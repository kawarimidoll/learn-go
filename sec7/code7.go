package main

import (
	"fmt"
	"time"
)

func goroutine1(ch chan string) {
	for {
		ch <- "packet from 1"
		time.Sleep(1000 * time.Millisecond)
	}
}

func goroutine2(ch chan string) {
	for {
		ch <- "packet from 2"
		time.Sleep(2000 * time.Millisecond)
	}
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go goroutine1(c1)
	go goroutine2(c2)

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}
