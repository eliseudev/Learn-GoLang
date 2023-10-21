package main

import "fmt"

func main() {

	fmt.Printf("Anônimas...")

	func() {
		fmt.Println("Olá Mundo...")
	}()

	func(texto string) {
		fmt.Println(texto)
	}("Olá parametro...")

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s %d", texto, 10)
	}("Passando parâmetro")

	fmt.Println(retorno)
}
