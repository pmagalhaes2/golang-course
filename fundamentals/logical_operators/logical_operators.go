package main

import "fmt"

func comprar(trab1, trab2 bool) (bool, bool, bool) {
	comprarTv50 := trab1 && trab2
	comprarTv32 := trab1 != trab2
	comprarSorvete := trab1 || trab2

	return comprarTv50, comprarTv32, comprarSorvete
}

func main() {
	tv50, tv32, sorvete := comprar(false, false)
	fmt.Printf("TV50: %t, TV32: %t, Sorvete: %t, Saudável: %t", tv50, tv32, sorvete, !sorvete)
}
