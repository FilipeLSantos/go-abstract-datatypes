# FlexEntregas (Go) - Estrutura do Projeto

Este repositório contem a estrutura base para o cenario de simulacao de entregas com diferentes meios, usando abstracao para esconder detalhes internos.

## Objetivo da estrutura

- Separar dominio, aplicacao e infraestrutura.
- Permitir troca de meio de entrega sem alterar o fluxo principal.
- Facilitar adicao de novos meios com baixo acoplamento.

## Arvore de pastas

```text
.
├── cmd/
│   └── flexentregas/
├── docs/
├── internal/
│   ├── application/
│   │   └── simulacao/
│   ├── domain/
│   │   ├── entrega/
│   │   └── meio/
│   └── infrastructure/
│       └── meios/
│           ├── bicicleta/
│           ├── cavalo/
│           ├── drone/
│           └── moto/
├── test/
└── LICENSE
```

## Explicacao da estrutura

### Visao geral

A organizacao segue separacao por camadas para manter o codigo desacoplado:

- `domain`: regras de negocio e abstracoes centrais do problema.
- `application`: orquestracao dos casos de uso (simulacao).
- `infrastructure`: implementacoes concretas dos meios de entrega.
- `cmd`: ponto de entrada do programa (executavel).

Com isso, o fluxo principal da simulacao nao depende de detalhes internos de cada meio.

### Papel de cada pasta

- `cmd/flexentregas/`
	- Entrada da aplicacao.
	- Responsavel por iniciar a simulacao e compor as dependencias.

- `internal/domain/meio/`
	- Representa o tipo abstrato de meio de entrega.
	- Aqui ficam os contratos das operacoes obrigatorias:
		- `calcularTempo`
		- `calcularCusto`
		- `podeOperar`

- `internal/domain/entrega/`
	- Representa os dados da entrega.
	- Agrupa informacoes como distancia, regiao e clima.

- `internal/application/simulacao/`
	- Implementa o caso de uso de simular entrega.
	- Usa apenas a abstracao do meio, sem `if` especifico por tipo.

- `internal/infrastructure/meios/moto/`
	- Implementacao concreta do meio Moto.

- `internal/infrastructure/meios/bicicleta/`
	- Implementacao concreta do meio Bicicleta.

- `internal/infrastructure/meios/drone/`
	- Implementacao concreta do meio Drone.

- `internal/infrastructure/meios/cavalo/`
	- Implementacao concreta do meio Cavalo.

- `docs/`
	- Documentacao de arquitetura e cenarios de simulacao.

- `test/`
	- Espaco reservado para testes da simulacao e dos meios.

### Como essa estrutura atende ao enunciado

- Representar diferentes meios de entrega:
	- Cada meio fica isolado em seu proprio modulo em `infrastructure/meios`.

- Esconder detalhes internos de calculo:
	- O fluxo principal trabalha com a abstracao definida em `domain/meio`.

- Trocar o meio sem alterar a funcao principal:
	- A simulacao depende de contrato, nao de implementacao concreta.

- Adicionar novos meios no futuro com pouca alteracao:
	- Basta criar novo modulo concreto em `infrastructure/meios` obedecendo o contrato.

## Observacao

Este setup foi criado sem implementacao de regras em codigo, conforme solicitado.
