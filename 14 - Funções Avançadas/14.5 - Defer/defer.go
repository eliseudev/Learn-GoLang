package main

import "fmt"

func alunoAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada. Resultado será retornado!")
	fmt.Println("Entrando na função para saber se o aluno foi aprovado.")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	} else {
		return false
	}
}
func main() {
	fmt.Println("Resultado:", alunoAprovado(10, 20))
}
