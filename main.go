package main

import (
	"fmt"
)

// Conversão
type hotdog int

var b hotdog = 101

func main() {
	x := 10
	x = int(b)

	fmt.Printf("%T\n", x)

	fmt.Printf("%T\n", b)
}
