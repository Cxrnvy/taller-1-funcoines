package utilidades

import "fmt"

func ConvertirMoneda(dolares float64, moneda string) {
	if moneda == "Euros" || moneda == "euros" {
		fmt.Println(dolares, "USD equivalen a", dolares*0.92, "Euros")
	} else if moneda == "LB" || moneda == "lb" {
		fmt.Println(dolares, "USD equivalen a", dolares*0.79, "Libras")
	} else if moneda == "Won" || moneda == "won" {
		fmt.Println(dolares, "USD equivalen a", dolares*1350.50, "Won")
	} else if moneda == "BTC" || moneda == "btc" {
		fmt.Println(dolares, "USD equivalen a", dolares*0.000015, "BTC")
	} else {
		fmt.Println("Moneda no válida. Usa: euros, lb, won, btc.")
	}
}

func ContarVocales(palabra string) {

	a := 0
	e := 0
	i := 0
	o := 0
	u := 0

	for _, letra := range palabra {
		if letra == 'a' || letra == 'A' {
			a++
		} else if letra == 'e' || letra == 'E' {
			e++
		} else if letra == 'i' || letra == 'I' {
			i++
		} else if letra == 'o' || letra == 'O' {
			o++
		} else if letra == 'u' || letra == 'U' {
			u++
		}
	}

	fmt.Println("Resultados:")
	fmt.Println("A:", a)
	fmt.Println("E:", e)
	fmt.Println("I:", i)
	fmt.Println("O:", o)
	fmt.Println("U:", u)
}
