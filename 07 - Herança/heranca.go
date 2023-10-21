package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança...")

	pessoa1 := pessoa{"Eliseu", "Oliveira", 34, 169}
	fmt.Println(pessoa1)

	estudante1 := estudante{pessoa1, "Sistemas para internet", "Estácio"}
	fmt.Println(estudante1)

}
