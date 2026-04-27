package main

import "fmt"

type Engine interface {
	milesLest() uint8
}

type gasEngine struct { // definimos um tipo struct para o motor a gasolina
	mpg      uint8
	galloons uint8
}

type electricEngine struct {
	kwh   uint8
	mpkwh uint8
}

// Dessa forma definimos um método para o tipo gasEngine, assim temos
// acesso as propriedades do struct e podemos realizar operações com elas
// parece que basta receber o struct como parâmetro e a partir disso temos acesso as suas propriedades
// ai da pra chamar var gas = gasEngine{mpg: 10, galloons: 2}; gas.milesLest()
func (g gasEngine) milesLest() uint8 {
	return g.mpg * g.galloons
}

func (e electricEngine) milesLest() uint8 {
	return e.kwh * e.mpkwh
}

// AHH!! Se passar os () antes significa que é um método, se passar depois é uma função normal
func milesLest(g gasEngine) uint8 {
	return g.mpg * g.galloons
}

func canMakeIt(engine Engine, miles uint8) {
	if miles <= engine.milesLest() {
		fmt.Println("You can make it without stopping for fuel!")
	} else {
		fmt.Println("You only have", engine.milesLest(), "miles worth of fuel, better get gas!")
	}
}

func main() {
	var gas = gasEngine{mpg: 10, galloons: 2}
	fmt.Println(gas.milesLest())
	fmt.Println(milesLest(gas))

	var electric = electricEngine{kwh: 20, mpkwh: 2}
	fmt.Println(electric.milesLest())

	canMakeIt(gas, 30)
	canMakeIt(electric, 30)
}
