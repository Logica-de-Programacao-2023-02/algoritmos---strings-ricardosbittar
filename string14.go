package main

import (
	"fmt"
	"strconv"
)

func main() {
	var palavra string

	fmt.Println("Digite uma palavra:")
	fmt.Scanln(&palavra)
	if decrescente(palavra) {
		fmt.Println("A string é uma sequência numérica decrescente.")
	} else {
		fmt.Println("A string não é uma sequência numérica decrescente.")
	}
}

func decrescente(s string) bool {

	for i := 1; i < len(s); i++ {

		num1, err1 := strconv.Atoi(string(s[i-1]))
		num2, err2 := strconv.Atoi(string(s[i]))

		if err1 != nil || err2 != nil || num2 >= num1 {
			return false
		}
	}

	return true
}


