package main

import "fmt"

func main() {

	pessoa := Pessoa{
		Nome:   "Batata",
		Idade:  90,
		Altura: 1.34,
	}

	cachorro := Cachorro{
		Apelido: "Cenoura",
		Cor:     "Branca",
	}

	pessoa.Apresentacao()
	cachorro.Apresentacao()

	var apresentador Apresentador = &pessoa

	apresentador.Apresentacao()

	ReceberApresentador(apresentador)
}

type Pessoa struct {
	Nome   string
	Idade  int
	Altura float64
}

type Cachorro struct {
	Apelido string
	Cor     string
}

type Apresentador interface {
	Apresentacao()
}

func (pessoa *Pessoa) Apresentacao() {
	fmt.Printf("Olá meu nome é %v, tenho %v de idade e %v de altura.",
		pessoa.Nome, pessoa.Idade, pessoa.Altura)
}

func (pessoa *Cachorro) Apresentacao() {
	fmt.Printf("O nome do cachorro é %v, ele tem a cor %v.",
		pessoa.Apelido, pessoa.Cor)
}

func ReceberApresentador(alvo Apresentador) {
	fmt.Println("\nA pessoa alvo é ", alvo)
}
