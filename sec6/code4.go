package main

import "fmt"

// type assertion
func doSomething(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println(v + 2)
	case string:
		fmt.Println(v + "!")
	default:
		fmt.Println("I don't know")
	}
}

func main() {
	doSomething(10)
	doSomething("hey")
	doSomething(true)
}
