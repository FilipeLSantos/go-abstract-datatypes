package entrega

const (
	RegiaoUrbana     = "urbana"
	RegiaoRural      = "rural"
	RegiaoMontanhosa = "montanhosa"
)

const (
	ClimaSol        = "sol"
	ClimaChuva      = "chuva"
	ClimaVentoForte = "vento forte"
)

type Entrega struct {
	DistanciaKm float64
	Regiao      string
	Clima       string
}
