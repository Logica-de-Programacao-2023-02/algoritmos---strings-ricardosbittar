package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var p1, p2 string
	fmt.Println("Digite uma palavra:")
	fmt.Scanln(&p1)
	fmt.Println("Digite outra palavra:")
	fmt.Scanln(&p2)

	p1 = strings.ToUpper(strings.ReplaceAll(p1, " ", ""))
	p2 = strings.ToUpper(strings.ReplaceAll(p2, " ", ""))

	slice1 := strings.Split(p1, "")
	slice2 := strings.Split(p2, "")

	sort.Strings(slice1)
	sort.Strings(slice2)
	if strings.Join(slice1, "") == strings.Join(slice2, "") {
		fmt.Println("As palavras sao anagramas")
	} else {
		fmt.Println("As palavras nao sao anagramas")
	}
}
