package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// padrao multiplexador - consiste em juntar resultados de multiplos canais

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

// padrão generator
func titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)
			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}
	return c
}

func main() {
	c := juntar(
		titulo("https://www.amazon.com", "https://www.youtube.com"),
		titulo("https://www.github.com", "https://www.udemy.com"),
	)

	fmt.Println(<-c)
}
