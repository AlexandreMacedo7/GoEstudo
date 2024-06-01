package main

import (
	"fmt"
)

// Nivel 2
// Exercicio 4

func main() {

	a := 200
	fmt.Printf("%d, %b, %#x\n", a, a, a)

	b := a << 1
	fmt.Printf("%d, %b, %#x", b, b, b)

}
