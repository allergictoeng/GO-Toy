package main

import "fmt"

func main() {
	i := 1
	var p *int = nil

	p = &i // pegando o endereço da variável
	*p++   // desreferenciando ( obtendo o valor)
	i++
	fmt.Print(p, *p, i, &i)

	// go não tem aritmética de ponteiros
}
