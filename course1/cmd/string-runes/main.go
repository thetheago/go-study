package main

import (
	"fmt"
)

func main() {
	// var nome string = "résumé"
	// fmt.Printf("%c\n", nome[1])

	// exemplo acima da problema por causa do utf-8

	var nome []rune = []rune("résumé")
	fmt.Printf("%c\n", nome[1])
}
