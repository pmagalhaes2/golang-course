package main

import "math"

// Inicializando com letra maiuscula é PÚBLICO (visível dentro e fora do pacote)
// Inicializando com letra minúscula é PRIVADO (visível dentro apenas dentro do pacote)

// Ponto representa uma coordenado no plano cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cx = math.Abs(b.y - a.y)
	return
}

// Distância é responsável por calcular a distância linear entre dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
