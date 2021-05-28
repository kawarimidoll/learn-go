package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func multi(params ...int) (result int) {
	result = 1
	for i := 0; i < len(params); i++ {
		result *= params[i]
	}
	return
}

func acc(initial int) func() int {
	x := initial
	return func() int {
		x++
		return x
	}
}

func main() {
	ca := make([]int, 3)
	fmt.Printf("len=%d cap=%d value=%v\n", len(ca), cap(ca), ca)
	for i := 0; i < 3; i++ {
		ca = append(ca, i)
		fmt.Println(ca)
	}

	cb := make([]int, 0, 3)
	fmt.Printf("len=%d cap=%d value=%v\n", len(cb), cap(cb), cb)
	for i := 0; i < 3; i++ {
		cb = append(cb, i)
		fmt.Println(cb)
	}

	m := map[string]int{"apple": 10, "banana": 20}
	fmt.Println(m, m["apple"])
	fmt.Println(m["cherry"])
	m["cherry"] = 30
	fmt.Println(m["cherry"])

	fmt.Println(add(30, 20))
	fmt.Println(multi(30, 20))

	cnt := acc(1)
	fmt.Println(cnt())
	fmt.Println(cnt())
	fmt.Println(cnt())

	// exercise
	f := 1.11
	fmt.Println(int(f))

	n := map[string]int{"Mike": 20, "Nancy": 24, "Messi": 30}
	fmt.Printf("%T %v", n, n)
}
