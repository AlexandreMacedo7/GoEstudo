package main

import (
	"fmt"
	"time"
)

// Nivel 3
// Exercicio 4

func main() {

	anoemquenasci := 1996
	anoatual := time.Now().UTC().Year()

	for {
		if anoemquenasci == anoatual {
			break
		}
		fmt.Println(anoemquenasci)
		anoemquenasci++
	}
}
