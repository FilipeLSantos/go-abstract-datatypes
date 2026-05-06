package drone

import (
	"go-abstract-datatypes/internal/domain/entrega"
	"go-abstract-datatypes/internal/domain/meio"
)

func Criar() meio.MeioEntrega {
	return meio.MeioEntrega{
		Nome: "Drone",

		CalcularTempo: func(distancia float64) float64 {
			return distancia / 50.0
		},

		CalcularCusto: func(distancia float64) float64 {
			return distancia * 0.30
		},

		PodeOperar: func(entrega entrega.Entrega) bool {
			return entrega.Clima != "chuva" && entrega.Clima != "vento forte"
		},
	}
}
