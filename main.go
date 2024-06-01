package main

import (
	"fmt"
	"time"
)

// Nivel 3
// Exercicio 3

func main() {
	for i := 1996; i <= time.Now().UTC().Year(); i++ {
		fmt.Println(i)
	}
}
