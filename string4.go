package main

import "fmt"

func main() {
	var frase1, frase2 string

	fmt.Println("Digite a primeira frase: ")
	fmt.Scanln(&frase1)
	fmt.Println("Digite a segunda frase: ")
	fmt.Scanln(&frase2)

	igual := true

	if frase1 != frase2 {
		igual = false
	}

	if igual {
		fmt.Println("As duas frases sao identicas")
	} else {
		fmt.Println("As duas frases nao sao identicas")
	}
}
