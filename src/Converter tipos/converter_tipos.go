package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Convertendo de inteiro para real
	var inteiro int = 55
	var realDoInteiro = float64(inteiro)

	fmt.Printf("Primeira conversão:\n\nInteiro: %d\nReal: %.2f", inteiro, realDoInteiro)

	// Mudando o valor do float e convertendo novamente para inteiro
	realDoInteiro = 879.00
	inteiro = int(realDoInteiro)

	fmt.Printf("\n\nSegunda conversão:\n\nReal: %.2f\nInteiro: %d", realDoInteiro, inteiro)

	numeroEmString := "154"
	inteiroDaString, erro := strconv.Atoi(numeroEmString)

	// A função Atoi (alfanumérico para inteiro) retorna o valor e um possível erro
	// que se for diferente de nil, existe.
	// Exemplo de verificação
	//
	// inteiroDaString, erro := strconv.Atoi(numeroEmString)
	//
	// if erro == nil {
	// 	fmt.Println("\nSucesso na conversão")
	// } else {
	// 	fmt.Println("Erro na conversão")
	// }

	fmt.Printf("\n\nTerceira conversão:\n\nNúmero em string: %s\nInteiro: %d\n", numeroEmString, inteiroDaString)
	fmt.Println("Erro: ", erro)
}
