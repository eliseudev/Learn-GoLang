package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("A área da forma é: %0.2f \n", f.area())
}

type retangulo struct {
	largura float64
	altura  float64
}

type circulo struct {
	raio float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func main() {
	raio := retangulo{10, 15}
	escreverArea(raio)

	c := circulo{10}
	escreverArea(c)
}
