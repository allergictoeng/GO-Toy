package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	fmt.Println("nova")
	fmt.Println("linha.")
	x := 3.1415
	//fmt.Println(" valor de X" + x) isso não funciona
	xs := fmt.Sprint(x)
	fmt.Println("valor de x" + xs)
	fmt.Println("O valor é", x)
	fmt.Printf("O valor de x é %.2f", x)
	a := 1
	b := 1.99999
	c := false
	d := "olá"
	fmt.Printf("\n %d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n %v %v %v %v", a, b, c, d) //mais genérico
}
