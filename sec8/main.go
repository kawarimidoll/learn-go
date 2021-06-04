package main

import (
	"fmt"

	"mymodule/mylib"
	"mymodule/mylib/sub"
)

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(mylib.Average(s))
	sub.Hello()
}
