package main

import "fmt"

type myError struct {
	message string
}

func (e *myError) Error() string {
	return fmt.Sprintf("Error: %s", e.message)
}

func validateSalary(salary int) error {
	if salary < 15000 {
		return &myError{message: "el salario ingresado no alcanza el mínimo imponible"}
	}
	return nil
}

func main () {
	var salary int = 12000

	err := validateSalary(salary)

	if err != nil {
		fmt.Println(err)
		return 
	}
	
	fmt.Println("Debe pagar impuestos")
}