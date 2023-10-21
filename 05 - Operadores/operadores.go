package main

import "fmt"

func main() {
	//Aritimericos
	soma := 10 + 10
	subtracao := 10 - 10
	multiplicacao := 10 * 5
	divisao := 10 / 4
	restoDivisao := 10 % 2

	fmt.Println(soma, subtracao, multiplicacao, divisao, restoDivisao)

	//Relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)

	//Atribuição
	var str string = "String"
	str1 := "string"
	fmt.Println(str, str1)

	//Lógicos
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)

	var texto string
	if 10 > 5 {
		texto = "é maior que 5"
	} else {
		texto = "é menor que 5"
	}
	fmt.Println(texto)

}
