package main

import (
	"fmt"
	"strings"
)

func main() {
	var p1, p2 string

	fmt.Println("Digite uma palavra:")
	fmt.Scanln(&p1)
	fmt.Println("Digite outra palavra:")
	fmt.Scanln(&p2)
	if strings.Contains(p1, p2) {
		fmt.Println("A segunda palavra esta dentro da primeira")
	} else {
		fmt.Println("A segunda palavra nao esta dentro da primeira ")
	}
}
