package main

import (
	"fmt"
	"time"
)

// channel (canal) - É a forma de comunicação entre as goroutines
// channel é um tipo

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base // enviando dados

	time.Sleep(time.Second) // esse tempo é só pra ver acontecendo o processo
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base
}

func main() {
	c := make(chan int)
	go doisTresQuatroVezes(2, c)

	a, b := <-c, <-c // recebendo os dados do canal (ele aguarda os dois valores chegarem pra continuar)
	fmt.Println(a, b)
	fmt.Println(<-c)
}
