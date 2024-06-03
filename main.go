package main

import (
	"fmt"
)

var x [5]int
var y [5]int

func main() {

	x[0] = 1
	y[2] = 1

	fmt.Println(x)
	fmt.Println(y)

}
