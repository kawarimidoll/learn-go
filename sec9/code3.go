package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println(os.Getwd())
	content, err := os.ReadFile("code3.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))
	if err := os.WriteFile("io_writed.go", content, 0666); err != nil {
		log.Fatal(err)
	}

	r := bytes.NewBuffer([]byte("abcd"))
	data, _ := io.ReadAll(r)
	fmt.Println(string(data))
}
