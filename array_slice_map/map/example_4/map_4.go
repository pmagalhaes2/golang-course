package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela S.": 1234.50,
			"Guga P.":     8456.78,
		},
		"J": {
			"José J.": 1154.66,
		},
		"P": {
			"Paulo M.": 4500.00,
		},
	}

	delete(funcsPorLetra, "J")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
