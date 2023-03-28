package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImpuestoSalario (t *testing.T) {
	t.Run("Salary minor to 50000", func(t *testing.T) {
		expextedTaxes := 0.0
		salary := 30000.0
		finalTaxes := impuestoSalario(salary)
		assert.Equal(t, expextedTaxes, finalTaxes, "Te result and the esc")
	})
	
	t.Run("Salary major to 50000", func(t *testing.T) {
		expextedTaxes := 9350.0
		salary := 55000.0
		finalTaxes := impuestoSalario(salary)
		assert.Equal(t, expextedTaxes, finalTaxes, "Te result and the esc")
	})

	t.Run("Salary mayor to 150000", func(t *testing.T) {
		expextedTaxes := 41850.0
		salary := 155000.0
		finalTaxes := impuestoSalario(salary)
		assert.Equal(t, expextedTaxes, finalTaxes, "Te result and the esc")
	})
} 