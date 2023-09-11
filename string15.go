package main

import (
	"fmt"
	"strings"
)

func main() {
	var palavra string

	fmt.Println("Digite uma palavra:")
	fmt.Scanln(&palavra)

	replacedString := trocarVogais(palavra)

	fmt.Println("String com vogais substituídas por asteriscos:", replacedString)
}

func trocarVogais(input string) string {

	replacer := strings.NewReplacer(
		"a", "*",
		"e", "*",
		"i", "*",
		"o", "*",
		"u", "*",
		"A", "*",
		"E", "*",
		"I", "*",
		"O", "*",
		"U", "*",
	)
	return replacer.Replace(input)
}
