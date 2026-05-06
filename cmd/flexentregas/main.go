package main

import (
	"fmt"
	"go-abstract-datatypes/internal/application/simulacao"
	"go-abstract-datatypes/internal/domain/entrega"
	"go-abstract-datatypes/internal/domain/meio"
	"go-abstract-datatypes/internal/infrastructure/meios/bicicleta"
	"go-abstract-datatypes/internal/infrastructure/meios/cavalo"
	"go-abstract-datatypes/internal/infrastructure/meios/drone"
	"go-abstract-datatypes/internal/infrastructure/meios/moto"
)

func main() {
	meios := []meio.MeioEntrega{
		moto.Criar(),
		bicicleta.Criar(),
		drone.Criar(),
		cavalo.Criar(),
	}

	entregas := []entrega.Entrega{
		{DistanciaKm: 10, Regiao: entrega.RegiaoUrbana, Clima: entrega.ClimaSol},
		{DistanciaKm: 20, Regiao: entrega.RegiaoRural, Clima: entrega.ClimaChuva},
		{DistanciaKm: 8, Regiao: entrega.RegiaoMontanhosa, Clima: entrega.ClimaVentoForte},
	}

	for i, entregaAtual := range entregas {
		fmt.Printf("\nEntrega %d -> distancia: %.2f km | regiao: %s | clima: %s\n", i+1, entregaAtual.DistanciaKm, entregaAtual.Regiao, entregaAtual.Clima)
		for _, meioEntrega := range meios {
			resultado := simulacao.SimularEntrega(meioEntrega, entregaAtual)
			fmt.Println(simulacao.FormatarResultado(resultado))
		}
	}
}
