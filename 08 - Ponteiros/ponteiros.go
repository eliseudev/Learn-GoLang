package main

import "fmt"

func main() {
	var variavel1 int8 = 10
	var variavel2 = variavel1
	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)

	//PONTEIRO É UM REFERÊNCIA DE MEMÓRIA
	var variavel3 int8
	var ponteiro *int8

	variavel3 = 100
	ponteiro = &variavel3

	fmt.Println(variavel3, ponteiro)
	//desreferenciação
	fmt.Println(variavel3, *ponteiro)

}
