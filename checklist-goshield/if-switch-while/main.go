package main

func main() {
	if true {
		println("This is an if statement")
	}

	switch "teste" {
	case "condition":
		println("This is a switch statement")
	case "teste":
		println("Vai cair aqui")
	default:
		println("This is the default case")
	}

	for {
		println("This is a while loop") // go não tem while
		break
	}
}
