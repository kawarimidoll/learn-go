package mylib

// ❯ go doc mylib Average

import (
	"fmt"
	"testing"
)

func Example() {
	v := Average([]int{1, 2, 3})
	fmt.Println(v)
}

func ExampleAverage() {
	v := Average([]int{1, 2, 3, 4, 5})
	fmt.Println(v)
}

func TestAverage(t *testing.T) {
	v := Average([]int{1, 2, 3, 4, 5})
	if v != 3 {
		t.Error("Expected 3, got", v)
	}
}
