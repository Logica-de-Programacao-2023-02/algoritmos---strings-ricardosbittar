package main

import (
	"fmt"
	"unicode"
)

func main() {
	var palavra string

	fmt.Println("Digite uma palavra:")
	fmt.Scanln(&palavra)

	isCamelCase, wordCount := checkCamelCase(palavra)

	if isCamelCase {
		fmt.Printf("A string está em camelCase e possui %d palavras.\n", wordCount)
	} else {
		fmt.Println("A string não está em camelCase.")
	}
}

func checkCamelCase(input string) (bool, int) {

	if len(input) == 0 {
		return false, 0
	}

	wordCount := 1
	isCamelCase := true


	for i := 0; i < len(input); i++ {

		if unicode.IsUpper(rune(input[i])) {

			if i > 0 {
				wordCount++
			}
		} else if !unicode.IsLower(rune(input[i])) {
			isCamelCase = false
			break
		}
	}

	return isCamelCase, wordCount
}
}
