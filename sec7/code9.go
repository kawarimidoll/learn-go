package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *Counter) Inc(key string) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.v[key]++
}

func (c *Counter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	m := make(map[string]int)
	// m["hey"] = 1
	// m["foo"] = 2

	go func() {
		for i := 0; i < 10; i++ {
			m["key"] += 1
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			m["key"] += 1
		}
	}()
	time.Sleep(1 * time.Second)
	fmt.Println(m)

	c := Counter{v: make(map[string]int)}
	go func() {
		for i := 0; i < 10; i++ {
			c.Inc("Key")
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			c.Inc("Key")
		}
	}()
	time.Sleep(1 * time.Second)
	fmt.Println(c.Value("Key"))
}
