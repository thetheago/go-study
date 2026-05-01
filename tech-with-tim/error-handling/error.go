package main

import "fmt"

func defferedFunc() {
	r := recover()

	if r == nil {
		fmt.Println("No error occurred")
	} else {
		fmt.Println("Error occurred:", r)
	}
}

func main() {
	defer defferedFunc()
	panic("This is a panic")
	fmt.Println("aa")
}
