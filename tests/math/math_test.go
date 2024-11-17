package math

import "testing"

const erroPadrao = "Valor esperado %v, mas o resultado encontrado foi %v."

// sempre deve seguir a convenção de prefixo Test antes de cada teste
func TestMedia(t *testing.T) {
	valorEsperado := 7.28
	valor := Media(7.2, 9.9, 6.1, 5.9)

	if valor != valorEsperado {
		t.Errorf(erroPadrao, valorEsperado, valor)
	}
}
