package main

import "fmt"

func inc(numero int) {
	numero++
}

func inc2(numero *int) {
	*numero++
}

func main() {
	n := 1

	inc(n) // por valor
	fmt.Println(n)

	// & é usado para onter o endereço de uma variável
	inc2(&n)
	fmt.Println(n)
}
