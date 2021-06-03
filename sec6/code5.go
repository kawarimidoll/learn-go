package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// Stringer: like toString() in other languages
func (p Person) String() string {
	return fmt.Sprintf("My name is %v", p.Name)
}

type UserNotFound struct {
	Name string
}

// Error: custom error
// 同一エラーの比較のためポインタを使うべきとされている
func (u *UserNotFound) Error() string {
	return fmt.Sprintf("User not found: %v", u.Name)
}

func myFunc() error {
	ok := false
	if ok {
		return nil
	}
	return &UserNotFound{Name: "foo"}
}

func main() {
	mike := Person{"Mike", 10}
	fmt.Println(mike)

	if err := myFunc(); err != nil {
		fmt.Println(err)
	}

}
