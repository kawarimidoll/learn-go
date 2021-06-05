package main

import (
	"fmt"
	"regexp"
	"sort"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Format(time.RFC3339))
	fmt.Println(t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second())

	match, _ := regexp.MatchString("a([a-z]+)e", "apple")
	fmt.Println(match)

	r := regexp.MustCompile("^/(edit|save|view)/\\w+$")
	fmt.Println(r.FindString("/view/test"))
	fmt.Println(r.FindString("/edit/2"))

	nums := []int{1, 9, -3, -4, 2, 8, 0}
	fmt.Println(nums)
	sort.Ints(nums)
	fmt.Println(nums)
	sort.Slice(nums, func(i, j int) bool { return nums[i]*nums[i] < nums[j]*nums[j] })
	fmt.Println(nums)
}
