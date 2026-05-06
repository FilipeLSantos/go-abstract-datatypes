package bicicleta

import (
	"go-abstract-datatypes/internal/domain/entrega"
	"go-abstract-datatypes/internal/domain/meio"
)

func Criar() meio.MeioEntrega {
	return meio.MeioEntrega{
		Nome: "Bicicleta",
		CalcularTempo: func(distancia float64) float64 {
			return distancia / 12.0
		},
		CalcularCusto: func(distancia float64) float64 {
			return 0.0
		},
		PodeOperar: func(entrega entrega.Entrega) bool {
			return entrega.DistanciaKm <= 15.0
		},
	}
}
