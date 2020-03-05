package main

import "fmt"

func main() {
	arrayEstatico := [5]int{1, 2, 3, 4, 5}
	arrayDinamico := []int{1, 2, 3, 4, 5}

	for i := 6; i <= 20; i++ {
		arrayDinamico = append(arrayDinamico, i) // A função append retorna uma nova cópia do array
	}

	fmt.Println(arrayEstatico)
	fmt.Println(arrayDinamico)
}
