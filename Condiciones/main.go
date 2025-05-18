package main

import "fmt"

func main() {
	nombre := "Sebastian"
	edad := 19

	if nombre == "Sebastian" {
		fmt.Println(nombre)
	} else {
		fmt.Println("Hola desconocido")
	}

	if edad >= 18 {
		fmt.Println("Ya puedes votar")
	}

	if 8%2 == 0 {
		fmt.Println("El 8 par")
	} else {
		fmt.Println("El 8 impar")
	}

	if numero := 99; numero < 0 {
		fmt.Println("El numero es negativo")
	} else if numero < 10 {
		fmt.Println("El numero es menor que 10")
	} else {
		fmt.Println("El numero es mayor que 10")
	}
}
