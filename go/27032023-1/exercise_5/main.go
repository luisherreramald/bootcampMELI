/*
Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas. Por el momento solo tienen tarántulas, hamsters, perros y gatos, 
pero se espera que puedan darle refugio a muchos animales más.
Tienen los siguientes datos:
Perro 10 kg de alimento.
Gato 5 kg de alimento.
Hamster 250 g de alimento.
Tarántula 150 g de alimento.
Se solicita:
-Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal especificado y que retorne una función y 
un mensaje (en caso que no exista el animal).
-Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del tipo de animal especificado.

*/

package main

import (
	"errors"
	"fmt"
)

func main() {
	animalType := "hamstersss"
	animal, err := calcularAlimento(animalType)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("La cantidad de alimento para %s es de %f kg \n", animalType, float64(animal(10)) / float64(1000))
	}
}


func perroCantidad (cantidad int) int {
	return cantidad * 10000  
}

func gatoCantidad (cantidad int) int {
	return cantidad * 5000
}

func hamsterCantidad (cantidad int) int {
	return cantidad * 250 
}

func tarantulaCantidad (cantidad int) int {
	return cantidad * 150
}

func calcularAlimento(animal string) (func(cantidad int) int, error) {
	switch animal {
		case "perro": 
			return perroCantidad, nil
		case "gato":
			return gatoCantidad, nil
		case "hamster":
			return hamsterCantidad, nil
		case "tarantula":
			return tarantulaCantidad, nil
		default:
			return nil, errors.New("Ha ingresado una opción no válida")
	}
}
