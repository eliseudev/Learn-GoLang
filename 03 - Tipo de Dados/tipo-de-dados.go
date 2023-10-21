package main

import (
	"errors"
	"fmt"
)

func main() {

	//NUMERO INTEIROS
	var numero int = 100
	fmt.Println(numero)

	var numero2 uint = 1000
	fmt.Println(numero2)

	//alias
	//int32 = rune
	var numero3 rune = 10000
	fmt.Println(numero3)

	//byte = uint8
	var numero4 byte = 123
	fmt.Println(numero4)

	//FIM NUMERO INTEIROS

	//NUMERO REAIS
	var numeroReal float32 = 123.65
	fmt.Println(numeroReal)

	var numeroReal1 float64 = 321.67
	fmt.Println(numeroReal1)

	numeroReal3 := 12345.12
	fmt.Println(numeroReal3)

	//FIM NUMERO REAIS

	//STRING
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto 2"
	fmt.Println(str2)
	//FIM STRING

	//Erro
	var erro error = errors.New("Tipo erro")
	fmt.Println(erro)
}
