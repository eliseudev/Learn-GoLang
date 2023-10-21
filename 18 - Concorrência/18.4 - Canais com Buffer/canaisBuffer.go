package main

import "fmt"

func main() {
	canal := make(chan string)
	canal <- "Olá..."
	menssage := <-canal
	fmt.Println(menssage)
}
