package main

import "fmt"

func main() {
	slice := make([]int, 1000, 1000) // Slice de tamanho 1000 e capacidade 1000

	for i := 6; i <= 1000; i++ {
		if i%2 != 0 {
			slice = append(slice, i) // A função append retorna uma nova cópia do array
		}
	}

	fmt.Println(slice)
}
