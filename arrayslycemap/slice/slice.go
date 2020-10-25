package main

import "fmt"

func main() {
	//slice não é array!!!
	a1 := [3]int{1, 2, 3} // array (fixo)
	s1 := []int{1, 2, 3}  // slice (variavel)
	fmt.Println(a1, s1)

	a2 := [5]int{1, 2, 3, 4, 5} // slice são trechos de array
	s2 := a2[1:3]               // Slice não cria um novo array ele apenas cria um ponteiro que aponta pro primeiro elemento do array ultilizado no slice
	fmt.Println(a2, s2)

	s3 := a2[:2]
	fmt.Println(s3)

	//slice tem um tamanho e um ponteiro pra um elemento de um array
	s4 := s2[:1]
	fmt.Println(s4)
}
