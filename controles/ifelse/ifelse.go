package main

import "fmt"

func imprResultado(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado com nota", nota)
	} else {
		fmt.Println("Reprovado com nota", nota)
	}
}

func main() {
	imprResultado(7.3)
	imprResultado(5.1)
}
