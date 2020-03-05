package main

import "fmt"

func main() {
	// Tipos mais básicos
	var inteiro int = 42     // usar uint para o unsigned int
	var real float64 = 3.14  // também chamado de ponto flutuante
	var booleano bool = true // também pode ser false
	var texto string = "Esta é uma string"

	fmt.Println(inteiro, real, booleano, texto)

	var (
		inteiroVazio  int
		realVazio     float64
		booleanoVazio bool
		stringVazia   string
	)

	fmt.Printf("\nCaso não se atribua valor algum, o valor padrão será:\n\nInteiros: %d\nReais: %f\nBooleanos: %v\nStrings: %s (aspas vazias)", inteiroVazio, realVazio, booleanoVazio, stringVazia)
}
