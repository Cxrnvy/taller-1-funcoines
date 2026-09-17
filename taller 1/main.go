package main

import "fmt"

// Función para calcular el promedio de forma muy sencilla
func averageGrade(sumaNotas float64, cantidad int) float64 {
	return sumaNotas / float64(cantidad)
}

// Opción 1: Promedio de notas
func opcion1() {
	var cantidad int
	var nota, sumaNotas float64

	fmt.Print("Ingrese la cantidad de estudiantes: ")
	fmt.Scan(&cantidad)

	// Usamos un for para pedir cada nota y sumarla
	for i := 1; i <= cantidad; i++ {
		fmt.Printf("Ingrese la nota del estudiante %d (0 a 100): ", i)
		fmt.Scan(&nota)
		sumaNotas = sumaNotas + nota
	}

	// Calculamos el promedio usando nuestra función
	promedio := averageGrade(sumaNotas, cantidad)
	fmt.Printf("\nEl promedio del curso es: %.2f\n", promedio)

	// IF para determinar si es aprobado o reprobado
	if promedio >= 70 {
		fmt.Println("Estado: APROBADO")
	} else {
		fmt.Println("Estado: REPROBADO")
	}

	// SWITCH para mostrar el mensaje según el rango
	switch {
	case promedio >= 90 && promedio <= 100:
		fmt.Println("Mensaje: Excellent performance")
	case promedio >= 80 && promedio <= 89:
		fmt.Println("Mensaje: Good performance")
	case promedio >= 70 && promedio <= 79:
		fmt.Println("Mensaje: Satisfactory performance")
	case promedio < 70:
		fmt.Println("Mensaje: Needs improvement")
	}
}

// Opción 2: Suma de números del 1 al n
func opcion2() {
	var n, suma int
	fmt.Print("Ingrese el número límite (n): ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		suma = suma + i
	}
	fmt.Printf("La suma del 1 al %d es: %d\n", n, suma)
}

// Opción 3: Celsius a Fahrenheit
func opcion3() {
	var celsius float64
	fmt.Print("Ingrese la temperatura en Celsius: ")
	fmt.Scan(&celsius)

	fahrenheit := (celsius * 9 / 5) + 32
	fmt.Printf("%.2f Celsius son %.2f Fahrenheit\n", celsius, fahrenheit)
}

// Opción 4: Fahrenheit a Celsius
func opcion4() {
	var fahrenheit float64
	fmt.Print("Ingrese la temperatura en Fahrenheit: ")
	fmt.Scan(&fahrenheit)

	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%.2f Fahrenheit son %.2f Celsius\n", fahrenheit, celsius)
}

// Función principal
func main() {
	var opcion string

	for {
		fmt.Println("\n--- Menú Principal ---")
		fmt.Println("1. Promedio de estudiantes")
		fmt.Println("2. Suma de números del 1 al n")
		fmt.Println("3. Celsius a Fahrenheit")
		fmt.Println("4. Fahrenheit a Celsius")
		fmt.Println("0. Salir (o escriba 'salir')")
		fmt.Print("Seleccione una opción: ")

		fmt.Scan(&opcion)

		// Evaluamos directamente si es "0", "salir" o "Salir" sin importar librerías extra
		if opcion == "0" || opcion == "salir" || opcion == "Salir" {
			fmt.Println("Saliendo del programa...")
			break
		}

		// Ejecutamos la función correspondiente con un switch simple
		switch opcion {
		case "1":
			opcion1()
		case "2":
			opcion2()
		case "3":
			opcion3()
		case "4":
			opcion4()
		default:
			fmt.Println("Opción incorrecta, intente de nuevo.")
		}
	}
}
