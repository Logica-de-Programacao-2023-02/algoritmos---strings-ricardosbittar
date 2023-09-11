package main

import "fmt"

func main() {
	var palavra string

	fmt.Println("Digite uma palavra:")
	fmt.Scanln(&palavra)
	var fraseRev string

	fmt.Println("Digite uma frase:")
	fmt.Scanln(&palavra)

	for i := len(palavra) - 1; i >= 0; i-- {
		fraseRev += string(palavra[i])
		fmt.Println("A frase revertida é:", fraseRev)
	}
}
