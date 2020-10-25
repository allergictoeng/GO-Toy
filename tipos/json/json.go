package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	// struct -> json
	p1 := produto{1, "Note", 1899.0, []string{"Promo", "Blabla"}}
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	// json -> struct
	var p2 produto
	jsonString := `{"id":8,"nome":"sei la","preco":30,"tags":["pitanga","fofinha"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2.Tags[1])
}
