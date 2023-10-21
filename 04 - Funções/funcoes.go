package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculoMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	sub := n1 - n2
	return soma, sub
}

func main() {
	somar := somar(10, 20)
	fmt.Println(somar)

	var f = func(txt string) {
		fmt.Println(txt)
	}

	f("Texto da função 2")

	resultadoSoma, resultadoSub := calculoMatematicos(20, 30)
	fmt.Println(resultadoSoma, resultadoSub)
	//Retornar somente um resultado de uma função com dois retornos
	resultado1, _ := calculoMatematicos(10, 15)
	fmt.Println(resultado1)

}
