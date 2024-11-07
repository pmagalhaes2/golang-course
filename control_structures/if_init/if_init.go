package main

import (
	"fmt"
	"math/rand"
	"time"
)

func geraNumeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	if i := geraNumeroAleatorio(); i > 5 { // sintaxe também suportada no switch
		fmt.Println("Ganhou!")
	} else {
		fmt.Println("Perdeu!")
	}
}