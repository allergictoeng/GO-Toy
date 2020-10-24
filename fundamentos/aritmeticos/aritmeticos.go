package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2
	fmt.Println("Soma =>", a+b)
	fmt.Println("SUb =>", a-b)
	fmt.Println("To com preguiça de fazer o basico e vou pular")

	//bitwise
	fmt.Println("E =>", a&b)
	fmt.Println("ou =>", a|b)
	fmt.Println("Xor =>", a^b)

	c := 3.0
	d := 2.0
	// lib math
	fmt.Println("Maior => ", math.Max(float64(a), float64(b)))
	fmt.Println("Menor =>", math.Min(c, d))
}
