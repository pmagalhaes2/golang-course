package main

import "fmt"

func media(numeros ...float64) float64 {
	total := 0.0
	for _, numero := range numeros {
		total += numero
	}

	return total / float64(len(numeros))
}

func main() {
	fmt.Printf("Média: %.2f", media(2.0, 3.0, 4.0))
}
