package main

import (
	"fmt"
	"strings"
)

func main() {
	var frase string

	fmt.Println("Digite uma frase:")
	fmt.Scanln(&frase)
	tirarEsp := strings.ReplaceAll(frase, " ", "")

	fmt.Println("A sua frase sem espacos fica:", tirarEsp)

}
