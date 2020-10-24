package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	// unsigned uint8,16,32,64
	var b byte = 255
	fmt.Println(reflect.TypeOf(b))

	// signed int8,16,32,64
	i1 := math.MaxInt64
	fmt.Println("O valor maximo do int é", i1)
	fmt.Println(reflect.TypeOf(i1))

	var i2 rune = 'a' // mapeamento da tabela Unicode (int32)
	fmt.Println("rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	// fiquei com preguiça de colocar o código int64 e boolean
}
