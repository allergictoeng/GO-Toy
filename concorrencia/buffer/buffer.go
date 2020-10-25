package main

import "fmt"

func rotina(ch chan int) {
	fmt.Println("Executou")
	ch <- 1
	ch <- 2
	ch <- 3 // só vem até aqui pois o tamanho do buffer é 3
	ch <- 4 // os outros só entram no buffer de acordo com os que estão são consumidos
	ch <- 5
	ch <- 6
}

func main() {
	ch := make(chan int, 3) // buffer tamanho 3
	go rotina(ch)
	fmt.Println(<-ch)
}
