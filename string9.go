package main

import (
	"fmt"
	"strings"
)

func main() {
	var frase, antigo, novo string

	fmt.Println("Digitie uma frase:")
	fmt.Scanln(&frase)
	fmt.Println("Digitie uma letra que deve ser substituida:")
	fmt.Scanln(&antigo)
	fmt.Println("Digitie uma letra que deve subsitituir a letra:")
	fmt.Scanln(&novo)

	novaString := strings.ReplaceAll(frase, antigo, novo)
	fmt.Println("A frase com as letras substituidas é:", novaString)
}
