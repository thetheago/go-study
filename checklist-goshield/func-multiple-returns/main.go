package main

import "fmt"

func functionThatReturnIntWithSucess() (int, error) {
	return 1, nil
}

func functionThatReturnError() (int, error) {
	return 0, fmt.Errorf("An error occurred")
}

func main() {
	fmt.Println("Hello, World!")

	number, err := functionThatReturnIntWithSucess()

	if err != nil {
		fmt.Println("Error occurred:", err)
	} else {
		fmt.Println("Function that return int with success: ", number)
	}

	number, err = functionThatReturnError()

	if err != nil {
		fmt.Println("Error occurred: ", err)
	} else {
		fmt.Println("No error, the number is: ", number)
	}

}
