package main

import (
	"fmt"
)

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

// Método: função com receiver (receptor)
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	prod := produto{
		nome:     "Lápis",
		preco:    8.00,
		desconto: 0.05,
	}

	fmt.Printf("Produto %s - Valor: R$ %.2f - Valor após desconto: R$ %.2f", prod.nome, prod.preco, prod.precoComDesconto())
}
