package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	UserId    int
	Id        int
	Title     string
	Completed bool
}

func main() {
	res, _ := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	defer res.Body.Close()

	u := User{}
	body, _ := io.ReadAll(res.Body)
	if err := json.Unmarshal(body, &u); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%#v\n", u)

	user, _ := json.Marshal(u)
	fmt.Printf(string(user))
}
