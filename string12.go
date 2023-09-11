package main

import (
	"fmt"
	"strings"
)

func main() {
	var palavra string

	fmt.Println("Digite uma palavra:")
	fmt.Scanln(&palavra)

	palavra = strings.ToUpper(strings.ReplaceAll(palavra, " ", ""))

	if palindromo(palavra) {
		fmt.Println("A palavra é um Palindromo")
	} else {
		fmt.Println("A palavra nao é um Palindromo")

	}

}
func palindromo(x string) bool {

	for i := 0; i < len(x)/2; i++ {
		if x[i] != x[len(x)]-1-i {
			return false
		}
		return true
	}
}