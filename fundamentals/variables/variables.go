package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2

	area := PI * math.Pow(raio, 2)

	fmt.Printf("A área do círculo é %.2f\n", area)

	g, h, i := 2, false, "aaa"

	fmt.Println(g, h, i)
}