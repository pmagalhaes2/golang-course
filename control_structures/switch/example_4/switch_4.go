package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "texto"
	case func():
		return "função"
	case bool:
		return "booleano"
	default:
		return "inválido"
	}
}

func main() {
	fmt.Println(tipo(2))
	fmt.Println(tipo(2.5))
	fmt.Println(tipo("Teste"))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(true))
	fmt.Println(tipo(time.Now()))
}
