package main

import (
	"fmt"
	"log"
	"strings"
)

// Bar this just for test
func Bar() string {
	return "good"
}

func main() {
	a := 1
	if a == 1 {
		log.Fatal("test err")
	}
	fmt.Println("vim-go")
	fmt.Println("hello world")
	fmt.Println(strings.ToUpper("good"))
}
