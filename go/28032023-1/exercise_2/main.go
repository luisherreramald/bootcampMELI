/*

Una empresa necesita realizar una buena gestión de sus empleados, para esto realizaremos un pequeño programa nos ayudará a gestionar correctamente dichos empleados. 
Los objetivos son:
Crear una estructura Person con los campos ID, Name, DateOfBirth.
Crear una estructura Employee con los campos: ID, Position y una composicion con la estructura Person.
Realizar el método a la estructura Employe que se llame PrintEmployee(), lo que hará es realizar la impresión de los campos de un empleado.
Instanciar en la función main() tanto una Person como un Employee cargando sus respectivos campos y por último ejecutar el método PrintEmployee().

Si logras realizar este pequeño programa pudiste ayudar a la empresa a solucionar la gestión de los empleados.
*/

package main

import "fmt"

type Person struct {
	ID          int
	Name        string
	DateOfBirth string
}

type Employee struct {
	ID       int
	Position string
	Person
}

func (emp Employee) PrintEmployee() {
	fmt.Printf("Num Empleado: %d \nNombre: %s \nPosition: %s \nDateOfBirth: %s \n", emp.ID, emp.Name, emp.Position, emp.DateOfBirth)
}

func main () {
	emp1 := Employee{1, "Programador", Person{1, "Carlos Ramirez", "10-09-1987"}}
	emp1.PrintEmployee()
}