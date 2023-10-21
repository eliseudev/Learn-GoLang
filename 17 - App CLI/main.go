package main

import (
	"fmt"
	"linhaComando/app"
	"log"
	"os"
)

func main() {
	//go run main.go ip --host mercadolivre.com.br
	fmt.Println("Ponto de partida...")
	aplicacao := app.Gerar()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
