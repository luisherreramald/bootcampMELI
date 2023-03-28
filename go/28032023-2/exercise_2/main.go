/*
Algunas tiendas ecommerce necesitan realizar una funcionalidad en Go para administrar productos y retornar el valor del precio 
total. La empresa tiene 3 tipos de productos: Pequeño, Mediano y Grande. (Se espera que sean muchos más)

Y los costos adicionales son:
Pequeño: solo tiene el costo del producto
Mediano: el precio del producto + un 3% de mantenerlo en la tienda
Grande: el precio del producto + un 6% de mantenerlo en la tienda, y adicional a eso $2500 de costo de envío.

El porcentaje de mantenerlo en la tienda es en base al precio del producto.
El costo de mantener el producto en stock en la tienda es un porcentaje del precio del producto.

Se requiere una función factory que reciba el tipo de producto y el precio y 
retorne una interfaz Producto que tenga el método Precio.

Se debe poder ejecutar el método Precio y que el método me devuelva el precio total en base al costo del producto y 
los adicionales en caso que los tenga

*/

package main

import (
	"fmt"
	"errors"
)

const (
	pequeno = "pequeño"
	mediano = "mediano"
	grande  = "grande"
)

func main() {
	producto, err := Factory(grande, 10000)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Producto:", grande )
		fmt.Println("Precio:", producto.Precio())
	}
}

type Producto interface {
	Precio() float64
}

type ProductoPequeno struct {
	Costo float64
}

type ProductoMediano struct {
	Costo         float64
	Mantenimiento float64
}

type ProductoGrande struct {
	Costo         float64
	Mantenimiento float64
	CostoEnvío    float64
}

func (p ProductoPequeno) Precio() float64 {
	return p.Costo
}

func (p ProductoMediano) Precio() float64 {
	return p.Costo * (1 + p.Mantenimiento)
}

func (p ProductoGrande) Precio() float64 {
	return p.Costo * (1+p.Mantenimiento) + p.CostoEnvío
}

func Factory(productType string, costoProducto float64) (Producto, error) {

	switch productType {
		case pequeno:
			return ProductoPequeno { Costo: costoProducto }, nil
		case mediano:
			return ProductoMediano { Costo: costoProducto, Mantenimiento: 0.03}, nil
		case grande:
			return ProductoGrande { Costo: costoProducto, Mantenimiento: 0.06, CostoEnvío: 2500 }, nil
		default: return nil, errors.New("ProductType is not defined")
	}
}