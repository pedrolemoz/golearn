package struct

import "fmt"

type pessoa struct {
	nome  string
	idade int
	cpf   string
}

func main() {
	nome := "Pedro"
	idade := 20
	cpf := "000.000.000-00"

	pessoa1 := pessoa{
		nome:  nome,
		idade: idade,
		cpf:   cpf,
	}

	fmt.Printf("Pessoa 1\n\nNome: %s\nIdade: %d\nCPF: %s\n", pessoa1.nome, pessoa1.idade, pessoa1.cpf)
}
