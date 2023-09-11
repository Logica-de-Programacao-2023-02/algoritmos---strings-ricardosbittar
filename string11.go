package main

import (
	"fmt"
	"strings"
)


func removervogal(palavra string) string{

	vogais:= "aeiouAEIOU"
	semVogais:= strings.Map(func(r rune) rune) {
		if strings.ContainsRune(vogais, r){
			return -1

		}
		return r

	}, palavra)
return semVogais

}
func main() {
	var palavra string

	fmt.Println("Digite uma palavra:")
	fmt.Scanln(&palavra)

	semVogais:= removervogal(palavra)

	fmt.Println("A palavra sem vogais é:", semVogais)



}
//So consegui com ChatGPT