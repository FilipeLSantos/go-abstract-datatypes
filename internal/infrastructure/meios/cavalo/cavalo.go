package cavalo

import (
	"go-abstract-datatypes/internal/domain/entrega"
	"go-abstract-datatypes/internal/domain/meio"
)

func Criar() meio.MeioEntrega {
	return meio.MeioEntrega{
		Nome: "Cavalo",

		CalcularTempo: func(distancia float64) float64 {
			return distancia / 8.0
		},

		CalcularCusto: func(distancia float64) float64 {
			return distancia * 0.10
		},

		PodeOperar: func(entrega entrega.Entrega) bool {
			return entrega.Regiao == "rural"
		},
	}
}
