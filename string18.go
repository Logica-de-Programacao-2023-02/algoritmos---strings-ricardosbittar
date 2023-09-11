package main

import (
	"fmt"
	"unicode"
)

func main() {
	var palavra string

	fmt.Println("Digite uma palavra:")
	fmt.Scanln(&palavra)

	if numeros(palavra) {
		fmt.Println("A palavra contém apenas números.")
	} else {
		fmt.Println("A palavra não contém apenas números.")
	}
}

func numeros(x string) bool {
	// Percorre os caracteres da string
	for _, i := range x {
		// Verifica se o caractere não é um dígito numérico
		if !unicode.IsDigit(i) {
			return false
		}
	}
	return true
}
