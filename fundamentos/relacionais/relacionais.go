package main

import "fmt"

func main() {
	fmt.Println("String:", "lala" == "lala")
	fmt.Println(" os relacionais funcionam como no java, então nada de novo aqui!")

	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"Jo"}
	p2 := Pessoa{"Jo"}
	fmt.Println("Pessoas:", p1 == p2)
}
