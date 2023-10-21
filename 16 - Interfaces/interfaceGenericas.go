package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("Eliseu")
	generica(34)
	generica(1.69)
}
