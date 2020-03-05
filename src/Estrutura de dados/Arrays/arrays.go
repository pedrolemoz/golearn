package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}

	j := 0
	for i := 100; i <= 500; i += 100 {
		array[j] = i // Modificando os elementos de um array
		j++
	}

	fmt.Println(array)
}
