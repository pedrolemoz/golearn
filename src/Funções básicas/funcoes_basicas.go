package main

import (
	"fmt"
	"math"
)

func main() {
	primeiroNumero := 25
	segundoNumero := 5
	fmt.Println(primeiroNumero, "elevado à potência de", segundoNumero, "é:", int(elevar(float64(primeiroNumero), float64(segundoNumero))))
}

func elevar(x, y float64) float64 {
	return math.Pow(x, y)
}
