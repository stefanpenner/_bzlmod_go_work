package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func Output() (string, int) {
	return reverse.String("Hello"), reverse.Int(4321)
}

func main() {
	a, b := Output()
	fmt.Println(a, b)
}
