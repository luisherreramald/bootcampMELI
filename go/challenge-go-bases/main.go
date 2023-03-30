package main

import (
	"fmt"
	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {

	defer func () {
		err := recover()

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Ejecución finalizada")
	}()

	records, err := tickets.LoadTickets("tickets.csv")

	if err != nil {
		panic(err)
	}

	fmt.Printf("Total de vuelos: %d\n", len(records))
	total := tickets.GetTotalTickets("China", records)

	fmt.Printf("El total de vuelos a China es: %d\n", total)
	totalPassengers, _ := tickets.GetCountByPeriod("morning", records)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Total de pasajeros en la mañana: %d \n", totalPassengers)

	averageDestination, err := tickets.AverageDestination("China", records)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Proporcion de vuelos a China: %f\n", averageDestination)

}
