package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++ // Incremento de a 1
		/*Aqui se puede poner mucho texto*/
	}

	fmt.Println("-------------------------")

	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	fmt.Println("-------------------------")

	for rango := range 3 {
		fmt.Println("rango: ", rango)
	}

	for {
		fmt.Println("loop")
		break
	}

	fmt.Println("-------------------------")

	for valor := 0; valor < 10; valor++ {
		if valor%2 == 0 {
			continue
		}
		fmt.Printf("Este es numero impar: %d\n", valor)
	}
	var matrisKa [8][8]int
	var m, n int = 8, 8

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			fmt.Print(matrisKa[i][j], "\t")
		}
		fmt.Println()
	}

}
