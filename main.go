package main

import (
	"fmt"
)

// Criano os proprios tipos
type hotdog int

var b hotdog

func main() {

	fmt.Printf("%T\n", b)
}
