package main

import (
	"fmt"
)

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println("IntMain test - run with \"go test\"")
}
