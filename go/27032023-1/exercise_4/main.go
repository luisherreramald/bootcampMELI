/*
Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones de los/as estudiantes de un curso. 
Requieren calcular los valores mínimo, máximo y promedio de sus calificaciones.
Para eso, se solicita generar una función que indique qué tipo de cálculo se quiere realizar (mínimo, máximo o promedio) y 
que devuelva otra función y un mensaje (en caso que el cálculo no esté definido) que se le pueda pasar una cantidad N de enteros y 
devuelva el cálculo que se indicó en la función anterior.
*/

package main

import (
	"errors"
	"fmt"
)

const (
	minimo = "minimo"
	promedio = "promedio"
	maximo = "maximo"
)

func main() {
	operation, err := calcularEstadisticas(minimo)
	if err != nil {
		fmt.Println(err)
	 } else {
		fmt.Println(operation(0,1,2,3,4,5,6,7,8,9))
	}
}

func calcularMinimo(notas ...int) float64 {
	fmt.Println("------------MINIMO------------")
	var min = notas[0]

	for note:= 1; note < len(notas); note++ {
		if notas[note] < min {
			min = notas[note]
		}
	}

	return float64(min)
}

func calcularPromedio(notas ...int) float64 {
	fmt.Println("------------PROMEDIO------------")
	var suma  = 0

	for _, note := range notas {
		suma = suma + note
	}

	return float64(suma) / float64(len(notas))
}

func calcularMaximo(notas ...int) float64 {
	fmt.Println("------------MAXIMO------------")
	var max = notas[0]

	for note:= 1; note < len(notas); note++ {
		if notas[note] > max {
			max = notas[note]
		}
	}
	return float64(max)
}

func calcularEstadisticas(tipoFuncion string) (func(notas ...int) float64, error){

	switch tipoFuncion {
		case "minimo": 
			return calcularMinimo, nil
		case "maximo":
			return calcularMaximo, nil
		case "promedio":
			return calcularPromedio, nil
		default:
			return nil, errors.New("Ha ingresado una opción no válida")
	}
}