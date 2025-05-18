package main

import (
	"fmt"
)

func main() {
	var numero int = 10
	numero2 := 20
	fmt.Println(numero, numero2)

	var numeroEntero = 10
	var numeroDoble = 20.5
	resultado := float32(numeroEntero) + float32(numeroDoble)
	fmt.Println(resultado)

	var nombre1 string = "Sebastian"
	apellido := "Perez"
	nombreCompleto := fmt.Sprintf("%s %s", nombre1, apellido)
	fmt.Println(nombreCompleto)

	nombre := "Sebastian"
	var edad int = 17
	// saludo := nombre + " y su edad es " + strconv.Itoa(edad)
	saludo1 := fmt.Sprintf("%s y su edad es %d", nombre, edad)
	fmt.Println(saludo1)

}
