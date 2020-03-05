package main

import "fmt"

func main() {
	primeiroNumero := 5
	segundoNumero := 25
	efetuaOperacoes(primeiroNumero, segundoNumero)
}

func efetuaOperacoes(x, y int) {
	fmt.Println("Operações com", x, "e", y)
	fmt.Println("Adição:", x+y)
	fmt.Println("Subtração:", x-y)
	fmt.Println("Multiplicação:", x*y)
	fmt.Println("Divisão:", x+y)
	fmt.Println("Módulo da divisão:", x%y)
	fmt.Println("AND:", x&y)
	fmt.Println("OR:", x|y)
	fmt.Println("XOR:", x^y)
}
