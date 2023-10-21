package main

import (
	"fmt"
)

func main() {
	// i := 0
	// for i < 10 {
	// 	i++
	// 	fmt.Println("incrementando: ", i)
	// 	time.Sleep(time.Second)
	// }

	// for j := 0; j < 10; j++ {
	// 	fmt.Println("Incrementando J", j)
	// 	time.Sleep(time.Second)
	// }

	nome := []string{"Eliseu", "Allana", "Elias", "Laura"}

	for indice, valor := range nome {
		fmt.Println(indice, valor)
	}

	fmt.Println("===============")
	for _, valor := range nome {
		fmt.Println(valor)
	}

	fmt.Println("===============")
	for indice, valor := range "PALAVRA" {
		fmt.Println(indice, valor)
	}

	fmt.Println("===============")
	for indice, valor := range "PALAVRA" {
		fmt.Println(indice, string(valor))
	}

	fmt.Println("===============")
	usuario := map[string]string{
		"nome":      "Eliseu",
		"sobrenome": "Oliveira",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

}
