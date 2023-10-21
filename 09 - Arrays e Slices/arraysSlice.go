package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices...")

	var arrays [2]string
	arrays[0] = "Posição 0"
	fmt.Println(arrays)

	array1 := [2]string{"Eliseu", "Elias"}
	fmt.Println(array1)

	array2 := [...]int{1, 2, 3, 4}
	fmt.Println(array2)

	slice := []int{10, 20, 30, 40, 50, 60}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(arrays))

	slice = append(slice, 70)
	fmt.Println(slice)
}
