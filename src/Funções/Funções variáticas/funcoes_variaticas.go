package funcoes_variaticas

import "fmt"

func main() {
	fmt.Println(mediaElementos(7, 57, 6, 86, 24, 35, 76, 54)) // A função pode receber n argumentos
}

// Com a notação de ... type, podemos receber muitos valores sem precisar declará-los
// Os valores serão armazenados em uma espécie de slice

func mediaElementos(elementos ...int) int {
	var sum int

	for _, value := range elementos {
		sum += value
	}

	return sum / len(elementos)
}
