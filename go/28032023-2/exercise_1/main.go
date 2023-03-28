package main

/*
Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para imprimir el detalle de los datos 
de cada uno de ellos/as, de la siguiente manera:

Nombre: [Nombre del alumno]
Apellido: [Apellido del alumno]
DNI: [DNI del alumno]
Fecha: [Fecha ingreso alumno]

Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.
Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido, DNI, Fecha y que tenga un método 
detalle
*/

import "fmt"

type Alumnos struct {
	Nombre string
	Apellido string
	DNI int
	Fecha string
}

func (alumno Alumnos) detalle(){
	fmt.Printf("Nombre %s \nApellido %s \nDNI %d \nFecha %s", alumno.Nombre, alumno.Apellido, alumno.DNI, alumno.Fecha)
}

func main() {
	alumno := Alumnos{}
	alumno.Nombre = "Carlos Alberto"
	alumno.Apellido = "Hernandez Guevara"
	alumno.DNI = 12345334
	alumno.Fecha = "13-03-2023"

	alumno.detalle()
}