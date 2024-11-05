package main

import "fmt"

func somar(a int, b int) int {
	return a + b
}

func imprimir(v int) {
	fmt.Printf("O resultado da soma é igual a %v", v)
}

func main() {
	resultado := somar(3, 4)
	imprimir(resultado)
}
