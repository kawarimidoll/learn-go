package main

import "fmt"

type Vertex struct {
	X, Y int
}

func one_var(x int) {
	x = 1
}

func one_ptr(x *int) {
	*x = 1
}

func changeVertex1(v Vertex) {
	v.X = 10
}

func changeVertex2(v *Vertex) {
	v.X = 10
}

func main() {
	var n int = 100
	fmt.Println(n)

	fmt.Println(&n)

	var p *int = &n // p is pointer of n
	fmt.Println(p)
	fmt.Println(*p) // *p is the value pointed by p

	one_var(n)
	fmt.Println(n)
	one_ptr(&n)
	fmt.Println(n)

	var p2 *int = new(int)
	fmt.Println(p2)
	fmt.Println(*p2)
	var p3 *int
	fmt.Println(p3)
	// fmt.Println(*p3) panic

	// 'make' makes value
	s := make([]int, 0)
	fmt.Printf("%T %v\n", s, s)
	m := make(map[string]int)
	fmt.Printf("%T %v\n", m, m)
	c := make(chan int)
	fmt.Printf("%T %v\n", c, c)

	// 'new' makes pointer
	fmt.Printf("%T %v\n", p, p)
	t := new(struct{})
	fmt.Printf("%T %v\n", t, t)

	v := Vertex{X: 1, Y: 2}
	fmt.Printf("%T %v\n", v, v)
	v2 := Vertex{}
	fmt.Printf("%T %v\n", v2, v2)
	var v3 Vertex
	fmt.Printf("%T %v\n", v3, v3)
	v4 := new(Vertex)
	fmt.Printf("%T %v\n", v4, v4)
	v5 := &Vertex{}
	fmt.Printf("%T %v\n", v5, v5)

	changeVertex1(v)
	fmt.Println(v)
	changeVertex2(&v)
	fmt.Println(v)
}
