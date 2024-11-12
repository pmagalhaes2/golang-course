package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual int
}

func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
}

func main() {
	f := ferrari{"F40", false, 0}
	fmt.Println(f)
	f.ligarTurbo()

	// Criando uma variável do tipo interface esportivo, que é um ponteiro de ferrari
	var f2 esportivo = &ferrari{"F40", false, 0}
	f2.ligarTurbo()

	fmt.Println(f, f2)
}