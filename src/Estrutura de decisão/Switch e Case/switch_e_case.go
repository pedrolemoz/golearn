package main

import "fmt"

func main() {
	numero := 10000

	switch {
	case numero <= 10:
		fmt.Println("O número é menor ou igual a 10")
		break
	case numero <= 100:
		fmt.Println("O número é menor ou igual a 100")
		break
	case numero <= 1000:
		fmt.Println("O número é menor ou igual a 1000")
		break
	case numero <= 10000:
		fmt.Println("O número é menor ou igual a 10000")
		break
	case numero <= 100000:
		fmt.Println("O número é menor ou igual a 100000")
		break
	default:
		fmt.Println("O número é maior que 100000")
	}
}
