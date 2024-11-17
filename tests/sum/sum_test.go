package sum

import "testing"

func TestSoma(t *testing.T) {
	resultado := Soma(1.0, 3.0, 1.5, 2.5, 0.5, 0.5, 0.75, 0.25)
	valorEsperado := 10.0

	if resultado != valorEsperado {
		t.Errorf("Esperado %f, mas obteve %f", valorEsperado, resultado)
	}
}
