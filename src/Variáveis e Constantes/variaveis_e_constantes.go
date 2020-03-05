package main

import (
	"fmt"
	"math"
)

func main() {
	// Declarando com o tipo definido
	var nome string = "Pedro"
	// Declarando variáveis sem especificar o tipo (por inferência)
	var idade = 20
	// Declarando variáveis de forma simplificada (Go shorthand)
	altura := 1.78
	// Declarando constantes
	const valorDePi = math.Pi

	// Também é possível declarar de forma simplificada
	var (
		outraString = "Esta é outra string"
		outroNumero = 42
	)
	// Printando ao estilo do C
	fmt.Printf("O meu nome é %s, tenho %d anos e %.2f de altura.\nO valor de Pi é %.2f.", nome, idade, altura, valorDePi)
	// Printando sem quebra de linha
	fmt.Print("\nOutra string: ", outraString, "\nOutro número: ", outroNumero)
}
