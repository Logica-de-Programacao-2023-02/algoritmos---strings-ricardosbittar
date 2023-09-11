package main

import (
	"fmt"
	"strconv"
)

func main() {
	var frase string
	fmt.Println("Escreva uma frase:")
	fmt.Scanln(&frase)

	num, err := strconv.ParseFloat(frase, 64)

	if err != nil {
		fmt.Printf("%v é um número float valido.\n", num)
	} else {
		fmt.Printf("%v nao é um float valido")
	}
}
