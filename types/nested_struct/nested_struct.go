package main

import "fmt"

type item struct {
	produtoID int
	qtd       int
	preco     float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0

	for _, item := range p.itens {
		total += item.preco * float64(item.qtd)
	}
	return total
}

func main() {
	pedidos := []pedido{
		{
			userID: 1,
			itens: []item{
				{produtoID: 1, qtd: 2, preco: 5.00},
				{produtoID: 2, qtd: 5, preco: 1.50},
			},
		},
		{
			userID: 2,
			itens: []item{
				{produtoID: 2, qtd: 5, preco: 1.50},
			},
		},
	}

	for _, p := range pedidos {
		fmt.Printf("Valor Total do Pedido do usuário %d: %.2f\n", p.userID, p.valorTotal())
	}
}
