package main

import "fmt"

func main() {
	var frase string
	var fraseRev string

	fmt.Println("Digite uma frase:")
	fmt.Scanln(&frase)

	for i := len(frase) - 1; i >= 0; i-- {
		fraseRev += string(frase[i])
		fmt.Println("A frase revertida é:", fraseRev)
	}
}
