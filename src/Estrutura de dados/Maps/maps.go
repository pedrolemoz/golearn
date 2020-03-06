package maps

import "fmt"

func main() {
	// Criamos maps exatamente como se fosse um slice

	novoMap := make(map[int]string)

	// O número representa um ID, e o nome, o valor que será armazenado
	// Assim como em Python e JavaScript, um map pode conter vários tipos
	// dentro dele, inclusive tipos definidos pelo usuário

	novoMap[12345] = "Pedro"
	novoMap[12346] = "Lucas"
	novoMap[12347] = "Marcos"

	// Se fosse em Python:
	// for k, v in range(novoMap)
	// Diferença apenas na mudança do in para o :=

	fmt.Printf("Valores iniciais\n\n")
	for key, value := range novoMap {
		fmt.Printf("ID: %d\tNome: %s\n\n", key, value)
	}

	// Deletando um valor

	delete(novoMap, 12347)

	fmt.Printf("Valores atualizados\n\n")
	for key, value := range novoMap {
		fmt.Printf("ID: %d\tNome: %s\n\n", key, value)
	}

	// Inicializando map com valores
	// Note que não foi utilizado o make
	// pois não estamos criando um map vazio

	identidades := map[string]string{
		"Pedro":  "123.456.789-00",
		"Carlos": "987.654.321-99",
		"João":   "789.456.123-09",
	}

	fmt.Printf("Lista de identidades\n\n")
	for key, value := range identidades {
		fmt.Printf("Nome: %s\nCPF: %s\n\n", key, value)
	}

}
