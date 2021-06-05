package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	res, _ := http.Get("http://example.com")
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	fmt.Println(string(body))
}
