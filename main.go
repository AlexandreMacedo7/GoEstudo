package main

import (
	"fmt"
)

//Exercicio 4

type tipo int

var x tipo

func main() {

	fmt.Printf("%v, %T\n", x, x)
	x = 40
	fmt.Printf("%v, %T", x, x)

}
