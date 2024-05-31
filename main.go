package main

import "fmt" //biblioteca padrão para E/S Formatada

func main() {

	x := "Oi "
	y := "de novo"

	//tipos de print

	//print normal - função retornar algo na tela
	fmt.Println("Olá")

	//printS - retorna uma String para função
	z := fmt.Sprint(x, y)

	fmt.Println(z)

	//Fprint writer interface

}
