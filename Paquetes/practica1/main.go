package main

import (
	"fmt"
	"practica1/operaciones"
	"practica1/saludo"
)

func main() {
	resultado := operaciones.Suma(5, 3)
	fmt.Println("La suma es:", resultado)

	mensaje := saludo.Saludar("Juan")
	fmt.Println(mensaje)
}
