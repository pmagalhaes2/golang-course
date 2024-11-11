package main

import "fmt"

func soma(a, b int) int {
	return a + b
}

func subtracao(a, b int) int {
	return a - b
}

func multiplicacao(a, b int) int {
	return a * b
}

func divisao(a, b int) int {
	return a / b
}

func exec(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}

func main() {
	resultado_soma := exec(soma, 10, 5)
	resultado_sub := exec(subtracao, 10, 5)
	resultado_mult := exec(multiplicacao, 10, 5)
	resultado_div := exec(divisao, 10, 5)

	fmt.Println(`Resultado soma: `, resultado_soma)
	fmt.Println(`Resultado subtração: `, resultado_sub)
	fmt.Println(`Resultado multiplicação: `, resultado_mult)
	fmt.Println(`Resultado divisão: `, resultado_div)
}
