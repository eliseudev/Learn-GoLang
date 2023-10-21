package main

import "fmt"

func main() {
	var variavel string = "variavel"
	fmt.Println(variavel)

	variavel1 := "Variavel 2"
	fmt.Println(variavel1)

	var (
		variavel3 string = "variavel 3"
		variavel4 string = "variavel 4"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "variavel 5", "variavel 6"
	fmt.Println(variavel5, variavel6)
}
