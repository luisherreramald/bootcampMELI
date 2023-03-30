package tickets

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)
type Ticket struct {
	ID string
	Nombre string
	Email string
	Pais string
	Hora string
	Precio string
}

func LoadTickets(nameFile string) ([]Ticket, error) {

	bookings := []Ticket{}

	file, err := os.Open(nameFile)
	
	if err != nil {
		return nil, fmt.Errorf("File dosent exist")
	}

	fileReader := csv.NewReader(file)
	for {
		record, err := fileReader.Read()

		if err != nil {
			if err == io.EOF {
				break
			}

			return nil, fmt.Errorf("File is broken")
		}

		booking := Ticket {
			ID: record[0],
			Nombre: record[1],
			Email: record[2],
			Pais: record[3],
			Hora: record[4],
			Precio: record[5],
		}

		bookings = append(bookings, booking)
	}
	
	return bookings, nil
}

func GetTotalTickets(destination string, bookings []Ticket) int {
	var counter = 0
	for _, booking := range bookings {
		if booking.Pais == destination {
			counter = counter + 1
		}
	}

	return counter
}

const (
	morning string = "morning"
	afternoon string = "afternoon"
	night string = "night"
	early_morning string = "early_morning"

)
// ejemplo 2
func GetCountByPeriod(time string, bookings []Ticket) (int, error) {
	totalPassengers := 0

	for _, booking := range bookings {
		hour := strings.Split(booking.Hora, ":")
		intHour, err := strconv.Atoi(hour[0])
		
		if err != nil {
			return 0, fmt.Errorf("Error at get hours")
		}

		switch time {
			case early_morning:
				if intHour >= 0 && intHour <= 6 {
					totalPassengers++
				}
			case morning:
				if intHour >= 7 && intHour <= 12 {
					totalPassengers++
				}
			case afternoon:
				if intHour >= 13 && intHour <= 19 {
					totalPassengers++
				}
			case night:
				if intHour >= 20 && intHour <= 23 {
					totalPassengers++
				}
		}
	}

	return totalPassengers, nil
}

// ejemplo 3
func AverageDestination(destination string, bookings []Ticket) (float64, error) {
	totalBookings := GetTotalTickets(destination, bookings)
	if totalBookings == 0 {
		return 0, fmt.Errorf("Error: No hay vuelos")
	}
	return float64(len(bookings)) / float64(totalBookings), nil
} 
