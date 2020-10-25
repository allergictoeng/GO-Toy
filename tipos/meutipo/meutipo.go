package main

type nota float64 // a parte relevante é esse alias

func (n nota) entre(inicio, fim float64) bool {
	return float64(n) >= inicio && float64(n) <= fim
}
func notaParaConceito(n nota) string {
	// preguiça de fazer o codigo daqui que é um
	// switch ou elseif
	return ""
}

func main() {
	// nothing to see
}
