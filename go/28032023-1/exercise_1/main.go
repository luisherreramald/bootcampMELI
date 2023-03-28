/*

Crear un programa que cumpla los siguiente puntos:
Tener una estructura llamada Product con los campos ID, Name, Price, Description y Category.   
Tener un slice global de Product llamado Products instanciado con valores.
2 métodos asociados a la estructura Product: Save(), GetAll(). El método Save() 
deberá tomar el slice de Products y añadir el producto desde el cual se llama al método. El método GetAll() deberá imprimir todos los productos guardados en el slice Products.
Una función getById() al cual se le deberá pasar un INT como parámetro y retorna el producto correspondiente al parámetro pasado. 
Ejecutar al menos una vez cada método y función definido desde main().

*/


package main

import (
	"errors"
	"fmt"
)


type Product struct {
	Id int
	Name string
	Price float64
	Description string
	Category string 
}

var Products = []Product{}

func (p Product) Save() {
	Products = append(Products, p)
}

func (p Product) GetAll() {
	for _, product := range Products {
		fmt.Printf("%d %s %f %s %s \n", product.Id, product.Name, product.Price, product.Description, product.Category)
	}
}

func getById (id int) (Product, error) {
	for _, product := range Products {
		if product.Id == id {
			return product, nil
		}
	}
	return Product{}, errors.New("Product not found")
}

func main () {

	p1 := Product{1, "Computadora", 2000.00, "Descripcion", "Electronicos"}
	p1.Save()
	fmt.Println("--------------Show Products-----------")
	p1.GetAll()
	producto, err := getById(3)
	fmt.Println("--------------Search Products-----------")
	if err == nil {
		fmt.Println(producto)
	} else {
		fmt.Println(err)
	}
}

