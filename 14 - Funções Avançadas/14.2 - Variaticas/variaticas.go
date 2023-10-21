package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	somas := soma(10, 20, 30, 40, 50, 60)
	fmt.Println("Total:", somas)

	escrever("Olá Mundão: ", 10, 20, 30, 40, 50, 60)
}
