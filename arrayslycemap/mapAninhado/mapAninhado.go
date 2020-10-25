package main

import "fmt"

func main() {

	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Guga": 12.3,
			"Gugu": 33.3,
		},
		"H": {
			"Homo": 1.2,
			"Homa": 1.3,
		},
		"P": {
			"Pu":  4.5,
			"Pru": 6.7,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
