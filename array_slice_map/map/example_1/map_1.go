package main

import "fmt"

func main() {
	// var aprovados map[int]string
	// mapas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[123456789] = "Maria"
	aprovados[123457862] = "José"
	aprovados[545458732] = "João"

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF : %d)\n", nome, cpf)
	}

	// excluir um dado de um map
	delete(aprovados, 545458732)
	fmt.Println(aprovados)
}
