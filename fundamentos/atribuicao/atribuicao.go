package main

import "fmt"

func main() {
	var b byte = 3
	fmt.Println(b)

	i := 3 // inferencia de tipo
	i += 3 // aditiva
	i -= 3 // ... repete a mesma logica com (/ * %)

	x, y := 1, 2
	fmt.Println(x, y)

	x, y = y, x
	fmt.Println(x, y)

}
