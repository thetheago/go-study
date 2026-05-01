package main

import (
	"fmt"
	"unsafe"
)

func async() {
	var num int8 = 2
	println("async", num)
}

func main() {
	var x int32 = 10

	sizeBytes := unsafe.Sizeof(x)
	sizeKB := float64(sizeBytes) / 1024

	fmt.Println("Bytes:", sizeBytes)
	fmt.Println("KB:", sizeKB)
}
