package main

import (
	"fmt"
)

const (
	a = iota // não tem um valor definitio, até ser utilizado, parecido como uma constante em numeros em sequencia
	b = iota
	c = iota
)

func main() {

	fmt.Print(a, b, c)

}
