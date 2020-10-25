package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteracao %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// ex: 1 (serial)
	//fale("Pitanga", "não late", 3)
	//fale("pitangão", "uauauaua", 1)

	// ex: 2 (paralelo) Não executa tudo por que a thread principal termina antes das outras
	//go fale("Pitanga", "hum!...", 500)
	//go fale("Pitangão", "auau", 500)
	//time.Sleep(time.Second * 5)
	//fmt.Println("Fim!")

	// ex: 3 (paralelo)
	go fale("Pitanga", "olha olha..", 10) // thread separada
	fale("Pitangão", "auauaua", 5)        // thread principal
}
