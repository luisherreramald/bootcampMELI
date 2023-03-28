/*
Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo, para cumplir el objetivo es necesario crear una 
función que devuelva el impuesto de un salario. 
Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17 % del sueldo y si gana más de $150.000 se le descontará además un 10 % (27% en total).
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(impuestoSalario(30000.00))
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
