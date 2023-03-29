/*
En tu función main, definí una variable llamada salary y asignale un valor de tipo “int”.
Creá un error personalizado con un struct que implemente Error() con el mensaje “Error: el salario es menor a 10.000" y lanzalo en caso de que “salary” sea menor o igual a  10.000. 
La validación debe ser hecha con la función “Is()” dentro del “main”.
*/

package main

import (
	"fmt"
	"errors"
)

type myError struct {
	message string
}

func (e *myError) Error() string {
	return fmt.Sprintf("Error: %s", e.message)
}

var ErrSalary = &myError{message: "el salario ingresado no alcanza el mínimo imponible"}

func validate(salary int) error {
	if salary < 10000 {
		return ErrSalary
	}
	return nil
}

func main () {
	var salary int = 11000
	err := validate(salary) 

	if err != nil {
		if errors.Is(err, ErrSalary) {
			fmt.Println(err)
			return 
		}
	} 

	fmt.Println("Debe pagar impuestos")
}