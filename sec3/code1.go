package main

import (
	"fmt"
	"os/user"
	"strings"
	"time"
)

const PI = 3.14159

func init() {
	fmt.Println("Init!")
}

func bazz() {
	fmt.Println("bazz!")
}

func main() {
	bazz()
	fmt.Println("Hello world!", "test!")
	fmt.Println("Hello", time.Now())
	fmt.Println(user.Current())

	var i int = 1
	f64 := 1.2
	var (
		s    string = "test"
		t, f bool   = true, false
	)
	fmt.Println(i, f64, s, t, f)

	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("10 - 1 =", 10-1)
	fmt.Println("3 * 2 =", 3*2)
	fmt.Println("10 / 2 =", 10/2)
	fmt.Println("10 / 3 =", 10/3)
	fmt.Println("10 % 2 =", 10%2)
	fmt.Println("10 % 3 =", 10%3)

	fmt.Println(strings.Replace(s, "t", "k", 1))

	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a)

	var b [2]int = [2]int{1, 2}
	fmt.Println(b)

	n := []int{1, 2, 3, 4, 5}
	fmt.Println(n, len(n), cap(n))
	fmt.Println(n[2])
	fmt.Println(n[2:4])
	fmt.Println(n[2:])
	fmt.Println(n[:2])
	n = append(n, 6, 7, 8)
	fmt.Println(n[:], len(n), cap(n)) // why cap is 10??
}
