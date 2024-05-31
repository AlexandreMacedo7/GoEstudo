package main

import (
	"fmt"
)

// Exercicio 5
// Tipo tem como subjacente int, isso permite que um tipo int receba uma conversão de Tipo
type tipo int

var x tipo
var y int

func main() {

	fmt.Printf("%v, %T\n", x, x)
	x = 40
	fmt.Printf("%v, %T\n", x, x)

	y = int(x)

	fmt.Printf("%v, %T", y, y)

}
