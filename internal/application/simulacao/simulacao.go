package simulacao

import (
	"fmt"
	"go-abstract-datatypes/internal/domain/entrega"
	"go-abstract-datatypes/internal/domain/meio"
)

type Resultado struct {
	Meio       string
	PodeOperar bool
	TempoHoras float64
	Custo      float64
}

func SimularEntrega(meioEntrega meio.MeioEntrega, entregaAtual entrega.Entrega) Resultado {
	resultado := Resultado{
		Meio: meioEntrega.Nome,
	}

	if !meioEntrega.PodeOperar(entregaAtual) {
		resultado.PodeOperar = false
		return resultado
	}

	resultado.PodeOperar = true
	resultado.TempoHoras = meioEntrega.CalcularTempo(entregaAtual.DistanciaKm)
	resultado.Custo = meioEntrega.CalcularCusto(entregaAtual.DistanciaKm)
	return resultado
}

func FormatarResultado(resultado Resultado) string {
	if !resultado.PodeOperar {
		return fmt.Sprintf("%s: este meio de entrega nao pode operar nessas condicoes.", resultado.Meio)
	}

	return fmt.Sprintf("%s: tempo estimado %.2f horas | custo estimado R$ %.2f", resultado.Meio, resultado.TempoHoras, resultado.Custo)
}
