package moto

import (
	"go-abstract-datatypes/internal/domain/entrega"
	"go-abstract-datatypes/internal/domain/meio"
)

func Criar() meio.MeioEntrega {
	return meio.MeioEntrega{
		Nome: "Moto",

		CalcularTempo: func(distancia float64) float64 {
			return distancia / 30.0
		},

		CalcularCusto: func(distancia float64) float64 {
			return distancia * 0.80
		},

		PodeOperar: func(entrega entrega.Entrega) bool {
			return entrega.Regiao != "montanhosa"
		},
	}
}
