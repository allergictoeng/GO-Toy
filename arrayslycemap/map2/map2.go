package main

import "fmt"

func main() {
	funcESalarios := map[string]float64{
		"pitanga":  12.3,
		"pitangão": 34.5,
	}

	funcESalarios["Pitanguetica"] = 99.0
	fmt.Println(funcESalarios)
	delete(funcESalarios, "inexistenteElenãoExcluiNAda")

	for nome, salario := range funcESalarios {
		fmt.Println(nome, salario)
	}
}
