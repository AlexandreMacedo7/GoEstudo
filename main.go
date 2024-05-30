package main

import "fmt" //biblioteca padrão para E/S Formatada

var variavel = "Variavel a nivel de pacote" // nivel de pacote

func main() {

	//Declaração de variavel - nivel de escopo
	x := 10
	y := "Olá bom dia"

	fmt.Printf("varivael: %v, %T\n", variavel, variavel)

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n\n", y, y)

	x = 20 //oporador de atribuição
	//x, z := 10, 5 //operador de declaração precisa ter pelo menos uma variavel nova

	fmt.Printf("x: %v, %T\n\n", x, x)

}
