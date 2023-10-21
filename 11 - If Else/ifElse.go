package main

import "fmt"

func main() {
	fmt.Println("Estrutura de controle")
	numero := 10

	if numero > 15 {
		fmt.Println("É maior que 15")
	} else {
		fmt.Println("menor ou igual a 15")
	}

	if outroNumero := numero; outroNumero > 10 {
		fmt.Printf("É maior que 10")
	} else if numero < -10 {
		fmt.Println("é menor que -10")
	} else {
		fmt.Printf("Entre 0 e -10")
	}
}
