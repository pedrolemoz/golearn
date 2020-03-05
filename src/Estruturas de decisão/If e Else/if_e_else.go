package main

import "fmt"

func main() {
	numero := 746

	if numero%2 == 0 {
		fmt.Println("O número é par")
	} else {
		fmt.Println("O número é ímpar")
	}

	// if e else também podem inicializar variáveis

	if outroNumero := 53; outroNumero%2 == 0 {
		fmt.Println("O número é par")
	} else {
		fmt.Println("O número é ímpar")
	}
}
