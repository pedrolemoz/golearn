package funcoes_basicas

import "fmt"

func main() {
	x := 547
	y := 12
	resultados := make([]int, 0) // Colocando 0 como tamanho, o slice crescerá dinamicamente
	resultados = append(resultados, soma(x, y))
	resultados = append(resultados, subtracao(x, y))
	resultados = append(resultados, multiplicacao(x, y))
	resultados = append(resultados, divisao(x, y))

	fmt.Printf("Operações entre %d e %d\n\nSoma: %d\nSubração: %d\nDivisão: %d\nDivisão: %d\n", x, y, resultados[0], resultados[1], resultados[2], resultados[3])

	max, min := getMax(x, y)

	fmt.Printf("Números primos de %d até %d\n\n", min, max)

	for i := min; i < max; i++ {
		if ehPrimo(i) {
			fmt.Printf("%d ", i)
		}
	}
}

// Funções de retorno simples, e que retornam um inteiro

func soma(x, y int) int {
	return x + y
}

func subtracao(x, y int) int {
	return x - y
}

func multiplicacao(x, y int) int {
	return x * y
}

func divisao(x, y int) int {
	return x / y
}

// Função que retorna um boolean

func ehPrimo(numero int) bool {
	counter := 0
	for i := 1; i < numero; i++ {
		if numero%i == 0 {
			counter++
		}

		if counter > 2 {
			break
		}
	}

	if counter < 2 {
		return true
	}
	return false

}

// Função que tem um múltiplo retorno

func getMax(x, y int) (int, int) {
	if x > y {
		return x, y
	}
	return y, x
}
