package main

import "fmt"

func main() {
	var arr [3]int32
	fmt.Println(&arr[0])
	fmt.Println(&arr[1])
	fmt.Println(&arr[2])
}
