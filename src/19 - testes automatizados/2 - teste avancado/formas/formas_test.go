package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {

	//TDD - Test Driven Development (NÃO SERÁ UTILIZADA NESTA APLICAÇÃO)

	t.Run("Retângulo", func(t *testing.T) {
		ret := Retangulo{10, 12}
		areaEsperada := float64(120)
		areaRecebida := ret.area()

		if areaEsperada != areaRecebida {
			t.Fatalf("A área recebida %f é diferente da área esperada %f", areaRecebida, areaEsperada)
		}
	})

	t.Run("Circulo", func(t *testing.T) {
		circ := Circulo{10}
		areaEsperada := float64(math.Pi * 100)
		areaRecebida := circ.area()

		if areaEsperada != areaRecebida {
			t.Fatalf("A área recebida %f é diferente da área esperada %f", areaRecebida, areaEsperada)
		}
	})

}
