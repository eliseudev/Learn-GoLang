package main

import "fmt"

func inverterSinal(numero int) int {
	return numero * -1
}

func inverterSinalPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	//Usando sem ponteiro, alterando o valor da copia da variavel.
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)

	//Usando ponteiro, alterando o valor do endereço da variavel.
	novoNumero := 40
	fmt.Println("Sem ponteiro:", novoNumero)
	inverterSinalPonteiro(&novoNumero)
	fmt.Println("Com ponteiro:", novoNumero)
}
