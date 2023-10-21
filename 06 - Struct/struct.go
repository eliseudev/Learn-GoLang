package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Printf("Structs...")

	var u usuario
	u.nome = "Elias"
	u.idade = 5
	fmt.Println(u)

	enderecoEx := endereco{"Rua Benjamin", 98}

	usuario2 := usuario{"Laura", 1, enderecoEx}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Eliseu"}
	fmt.Println(usuario3)
}
