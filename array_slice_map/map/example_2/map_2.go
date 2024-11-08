package main

import "fmt"

func main() {
	funcionarios := map[string]float64{
		"João B." : 11325.45,
		"Gabriela S.": 1500.20,
		"Pedro J.": 1200.00,
	}

	funcionarios["Rafaela F."] = 1350.00

	for nome, salario := range funcionarios {
		fmt.Printf("Funcionário(a): %s - Salário: %.2f\n", nome, salario)
	}
}