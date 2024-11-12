package main

import "fmt"

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

// interfaces são implementadas implicitamente
func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s = R$ %.2f", p.nome, p.preco)
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	pessoa1 := pessoa{"Maria", "José"}
	fmt.Println(pessoa1.toString())

	produto1 := produto{"Lápis", 5.50}
	fmt.Println(produto1.toString())

	produto2 := produto{"Calça Jeans", 159.90}
	imprimir(produto2)
}