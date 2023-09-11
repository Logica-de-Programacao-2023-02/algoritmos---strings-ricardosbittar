package main

import (
	"fmt"
	"strings"
)

func main() {
	var string string

	fmt.Println("Digite uma palavra qualquer:")
	fmt.Scanln(&string)

	maiusculo := strings.ToUpper(string)
	fmt.Println("A palavra com todos os caracteres em maiusculo fica:", maiusculo)
}
