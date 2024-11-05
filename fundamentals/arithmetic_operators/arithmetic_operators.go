package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Soma => ", a+b)
	fmt.Println("Subtração => ", a-b)
	fmt.Println("Divisão => ", a/b)
	fmt.Println("Multiplicação => ", a*b)
	fmt.Println("Módulo => ", a%b)

	c := 3.0
	d := 2.0

	// operação usando Math
	fmt.Println("Maior => ", math.Max(c, d))
	fmt.Println("Exponenciação =>", math.Pow(c, d))
}
