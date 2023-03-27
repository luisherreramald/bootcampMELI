package main

import (
	"fmt"
)

func main() {
	fmt.Println(impuestoSalario(150000.00))
}

func impuestoSalario(sueldo float64) float64 {
	if sueldo > 50000.00  && sueldo <= 150000.00 {
		return sueldo * 0.17
	}else if sueldo > 150000.00 {
		return sueldo * 0.27
	}else {
		return 0
	}
}
