package defer

import "fmt"

func main() {
	notas := make([]int, 100000, 100000)

	for i := 0; i < 100000; i++ {
		notas[i] = ((i*99999)*123456 + (7*8*9*10*11)/121314) % 10
	}

	fmt.Printf("Média dos valores: %d\n", calcularMedia(notas))
}

func calcularMedia(notas []int) int {
	defer fmt.Println("Fim da execução")
	var soma int
	for _, v := range notas {
		soma += v
	}

	return soma / len(notas)
}
