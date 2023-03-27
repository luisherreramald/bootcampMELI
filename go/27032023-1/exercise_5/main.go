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
