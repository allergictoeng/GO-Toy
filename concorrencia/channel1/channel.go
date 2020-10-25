package main

import "fmt"

func main() {
	ch := make(chan int, 1) // criando canal (canal é mecanismo de sincronismo)
	ch <- 1                 // enviando dados pro canal (escrita)
	<-ch                    // recebendo dados do canal (leitura)

	ch <- 2
	fmt.Println(<-ch)
}
