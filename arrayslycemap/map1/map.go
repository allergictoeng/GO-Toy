package main

import "fmt"

func main() {
	//var aprovados map[int]string

	aprovados := make(map[int]string)

	aprovados[123] = "Maria"
	aprovados[456] = "Pedro"
	aprovados[789] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}
	fmt.Printf(aprovados[123])
	delete(aprovados, 789)
	fmt.Println(aprovados[456])
}
