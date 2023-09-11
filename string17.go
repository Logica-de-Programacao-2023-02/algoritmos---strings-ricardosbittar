package main

import "fmt"

func main() {
	var palavra string

	fmt.Println("Digite uma palavra:")
	fmt.Scanln(&palavra)
	contarLetra := make(map[rune]int)

	for _, i := range palavra {
		if contarLetra[i] == 1 {
			fmt.Printf("%c", i)
		}
	}
	fmt.Println()
}
