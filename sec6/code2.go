package main

import "fmt"

type MyInt int

func (i MyInt) Double() int {
	// need to cast to int
	return int(i * 2)
}

func main() {
	myInt := MyInt(10)
	dbl := myInt.Double()
	fmt.Printf("%T %v\n", myInt, myInt)
	fmt.Printf("%T %v\n", dbl, dbl)
}
