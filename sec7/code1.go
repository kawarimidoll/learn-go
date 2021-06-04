package main

import (
	"fmt"
)

func normal(s string) {
	for i := 0; i < 5; i++ {
		// time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go normal("world")
	normal("hello")
	// time.Sleep(2000 * time.Millisecond)
}
