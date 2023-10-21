package main

import "fmt"

func recuperarFuncao() {
	if r := recover(); r != nil {
		fmt.Println("Execursão recuperada com sucesso!")
	}
}

func alunoAprovado(n1, n2 float32) bool {
	defer recuperarFuncao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A MÉDIA É EXATAMENTO 6!")
}

func main() {
	fmt.Println("Resultado:", alunoAprovado(6, 6))
	fmt.Println("Pós execursão.")
}
