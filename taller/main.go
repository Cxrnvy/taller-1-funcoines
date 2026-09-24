package main

import (
	"fmt"
	"taller/utilidades"
)

func main() {
	var opcion int

	for {
		fmt.Println("\n--- MENÚ ---")
		fmt.Println("1. Conversor de Monedas")
		fmt.Println("2. Contador de Vocales")
		fmt.Println("0. Salir")
		fmt.Print("Elige una opción: ")
		fmt.Scan(&opcion)

		// Condicionales (Tema 1.2) para el menú
		if opcion == 0 {
			fmt.Println("Fin del programa.")
			break
		} else if opcion == 1 {
			var dolares float64
			var moneda string

			fmt.Print("Ingresa dólares: ")
			fmt.Scan(&dolares)
			fmt.Print("Moneda (euros, lb, won, btc): ")
			fmt.Scan(&moneda)

			utilidades.ConvertirMoneda(dolares, moneda)

		} else if opcion == 2 {
			var palabra string

			fmt.Print("Escribe una sola palabra (sin espacios): ")
			fmt.Scan(&palabra)

			utilidades.ContarVocales(palabra)

		} else {
			fmt.Println("Opción no válida.")
		}
	}
}
