package main

import "testing"

func TestSoma(t *testing.T) {
	total := soma(5, 5)

	if total != 10 {
		t.Errorf("Resultado da soma é inválido: Resultado %d. Esperado: 10", total)
	}
}
