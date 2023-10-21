package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuário %s no banco de dados.\n", u.nome)
}

func (u usuario) maiorIdade() bool {
	return u.idade >= 18
}

//usando ponteiro
func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuarioDados := usuario{"Eliseu Oliveira", 34}
	fmt.Println(usuarioDados)

	usuarioDados.salvar()

	usuarioDados2 := usuarioDados.maiorIdade()

	fmt.Println(usuarioDados2)
	usuarioDados.fazerAniversario()
	fmt.Println(usuarioDados.idade)

}
