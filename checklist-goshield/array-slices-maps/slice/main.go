package main

import "fmt"

func modify(s []int) {
	s[0] = 999
}

func main() {
	// O slice é uma “visão” de um array, mas com tamanho dinâmico.
	// É basicamente tirar o tamanho do array, ja vira um slice

	s := []int{1, 2, 3}
	s = append(s, 4)

	fmt.Println(s)

	slice := []int{1, 2, 3}
	modify(slice)
	fmt.Println(slice) // mas muda, é um apontamento para array então é passado a referencia, não a copia
}
