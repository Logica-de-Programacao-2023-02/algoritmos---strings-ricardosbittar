package main

import (
	"fmt"
	"unicode"
)

func main() {
	var frase string
	fmt.Println("Digite uma frase:")
	fmt.Scanln(&frase)

	contemnum := false
	for _, i := range frase {
		if unicode.IsDigit(i) {
			contemnum = true
			break
		}

	}
	if contemnum {
		fmt.Println("A frase contem numeros")
	} else {
		fmt.Println("A frase nao contem numeros")
	}
}
