package main

import "fmt"

func main() {
	texto := "Texto simples"
	var ponteiro *string = &texto // O ponteiro do tipo string recebe o endereço de memória de texto
	fmt.Println("Conteúdo dentro do ponteiro:", *ponteiro, "\nEndereço de memória do ponteiro:", ponteiro)
	// Usando o ponteiro para alterar o conteúdo da variável
	*ponteiro = "Texto alterado"
	fmt.Println("\n\nConteúdo dentro do ponteiro:", *ponteiro, "\nEndereço de memória do ponteiro:", ponteiro)

}
