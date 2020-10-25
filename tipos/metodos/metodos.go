package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) nomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

func (p *pessoa) mudarNomeCompleto(nomeCompleto string) {
	partes := strings.Split(nomeCompleto, " ")
	p.nome = partes[0]
	p.sobrenome = partes[1]
}

func main() {
	p1 := pessoa{"Pedro", "Silva"}
	fmt.Println(p1.nomeCompleto())

	p1.mudarNomeCompleto("Pitanga Maria")
	fmt.Println(p1.nomeCompleto())
}
