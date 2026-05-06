package meio

import "go-abstract-datatypes/internal/domain/entrega"

type MeioEntrega struct {
	Nome          string
	CalcularTempo func(distancia float64) float64
	CalcularCusto func(distancia float64) float64
	PodeOperar    func(entrega entrega.Entrega) bool
}
