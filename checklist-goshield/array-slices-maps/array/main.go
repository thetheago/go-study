package main

import "fmt"

func modify(a [3]int) {
	a[0] = 999
}

func main() {
	// O array é um tipo mais rigido, tem tamanho fixo

	arr := [3]int{1, 2, 3}
	modify(arr)
	fmt.Println(arr) // [1 2 3] (não mudou)
}
