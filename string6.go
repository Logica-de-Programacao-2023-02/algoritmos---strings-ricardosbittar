package main

import "fmt"

func main() {
	var frase string
	fmt.Println("Digite uma frase:")
	fmt.Scanln(&frase)
	soma := 0

	for i := 0; i < len(frase); i++ {
		soma = soma + 1

	}
	fmt.Println("O número de caracteres desta frase é:", soma)
}
