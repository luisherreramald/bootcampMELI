package main

import (
	"fmt"
)

func main() {
	fmt.Println(calcularSalario(60 , "A"))

}

func convertirMinutosAHoras(minutos int) float64 {
	return float64(minutos / 60)
}

func calcularSalario (minutosMes int, categoria string) float64 {
	horas := convertirMinutosAHoras(minutosMes)
	switch categoria{
		case "A":
			return horas * 3000 + (horas * 3000 * 0.5)
		case "B":
			return horas * 1500 + (horas * 3000 * 0.2)
		case "C": 
			return horas * 1000

	}
	return 0
}