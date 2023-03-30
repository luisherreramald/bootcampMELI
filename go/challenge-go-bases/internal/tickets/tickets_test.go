package tickets

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var tickets, _ = LoadTickets("tickets.csv")

func TestGetTotalTickets(t *testing.T) {
	
	t.Run("Get total tickets to Brazil", func(t *testing.T){
		const totalTicketsExpected int = 45
		const country string = "Brazil"

		totalTicketsReal := GetTotalTickets(country, tickets)

		assert.Equal(t,totalTicketsExpected, totalTicketsReal)
	})

	t.Run("Get total tickets to China", func(t *testing.T){
		const totalTicketsExpected int = 178
		const country string = "China"

		totalTicketsReal := GetTotalTickets(country, tickets)

		assert.Equal(t,totalTicketsExpected, totalTicketsReal)
	})
}

func TestGetCountByPeriod(t *testing.T) {
	
	t.Run("Get total tickets at the morning", func(t *testing.T){
		const totalMorningTicketsExpected int = 151
		const time string = "night"

		totalMorningTicketsReal, _ := GetCountByPeriod(time, tickets)

		assert.Equal(t,totalMorningTicketsExpected, totalMorningTicketsReal)
	})

	t.Run("Get total tickets at the afternoon", func(t *testing.T){
		const totalAfternoonTicketsExpected int = 256
		const time string = "morning"

		totalAfternoonTicketsReal, _ := GetCountByPeriod(time, tickets)

		assert.Equal(t,totalAfternoonTicketsExpected, totalAfternoonTicketsReal)
	})
}

func TestAverageDestination(t *testing.T) {
	t.Run("Get brazil tickets proportion", func(t *testing.T){
		const totalProportionTicketsExpected float64 = 22.22222222222222
		const country string = "Brazil"

		totalProportionTicketsReal, _ := AverageDestination(country, tickets)

		assert.Equal(t,totalProportionTicketsExpected, totalProportionTicketsReal)
	})

	t.Run("Get China tickets proportion", func(t *testing.T){
		const totalProportionTicketsExpected float64 = 5.617977528089888
		const country string = "China"

		totalProportionTicketsReal, _ := AverageDestination(country, tickets)

		assert.Equal(t,totalProportionTicketsExpected, totalProportionTicketsReal)
	})

	t.Run("Get China tickets proportion", func(t *testing.T){
		const errMessage = "Error: No hay vuelos"
		const country string = "China"
		var emptyTickets = []Ticket{}

		_, err := AverageDestination(country, emptyTickets)
	
		assert.EqualError(t, err, errMessage)
	})
}
