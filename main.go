package main

import (
	"fmt"
)

// Nivel 3
// Exercicio 2

func main() {
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 1; j <= 3; j++ {
			fmt.Printf("\t %#U \n", i)
		}
	}

}
