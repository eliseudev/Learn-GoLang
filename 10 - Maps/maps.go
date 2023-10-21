package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps...")

	usuario := map[string]string{
		"nome":      "Eliseu",
		"sobrenome": "Oliveira",
	}
	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	usuario1 := map[int]string{
		1: "Elias",
		2: "Ataides",
		3: "Laura",
		4: "Ataides",
	}

	fmt.Println(usuario1[1], usuario1[3])

	//MAP ANINHADO
	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Eliseu",
			"segundo":  "Oliveira",
		},
		"curso": {
			"nomeCurso":   "Sistemas para internet",
			"instituicao": "Estácio",
		},
	}
	fmt.Println(usuario2["nome"], usuario2["curso"])

	usuario2["idioma"] = map[string]string{
		"nome": "Portugês",
	}

	fmt.Println(usuario2["nome"], usuario2["curso"], usuario2["idioma"])
}
