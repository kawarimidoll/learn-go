package main

import (
	"fmt"
	"time"
)

func routine1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	time.Sleep(2000 * time.Millisecond)
	c <- sum
}

func routine2(s []int, c chan int) {
	sum := 0
	for i := range s {
		sum += i
	}
	// time.Sleep(2000 * time.Millisecond)
	c <- sum
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int, 1)
	fmt.Println(s)
	fmt.Println(c)

	go routine1(s, c)
	go routine2(s, c)
	x := <-c
	y := <-c
	fmt.Println(x)
	fmt.Println(y)
}
