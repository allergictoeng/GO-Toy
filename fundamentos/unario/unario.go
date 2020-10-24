package main

import "fmt"

func main() {
	x := 1
	y := 2

	// apenas posfixada não tem prefixada
	x++
	fmt.Print(x)
	y--
	fmt.Print(y)
}
