package main

import (
	"fmt"
	"mudule/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Importando do arquivo main...")
	auxiliar.Escrever()
	erro := checkmail.ValidateFormat("eliseu.dev@outlook.com")

	if erro == nil {
		fmt.Println("Formato do email correto")
	} else {
		fmt.Println("Formato email incorreto.")
	}
}
