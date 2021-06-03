package main

import "fmt"

type Human interface {
	Say()
	Grow()
}

type Person struct {
	Name string
	Age  int
}

func (p Person) Say() {
	fmt.Println("I'm", p.Name, ",", p.Age, "years old.")
}

func (p *Person) Grow() {
	p.Age++
}

func main() {
	// Sayだけなら値定義でも使えるがGrowでポインタを使いたいので
	// アドレス定義を使う
	var mike Human = &Person{"Mike", 10}

	mike.Say()
	mike.Grow()
	mike.Say()
}
