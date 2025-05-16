package main

import (
	"fmt"
	"os"
)

func main() {
	envVar := os.Getenv("HOME")
	if envVar == "" {
		fmt.Println("HOME is not set")
	} else {
		fmt.Printf("HOME envairoment variable: %s\n", envVar)
	}

	file, err := os.Create("ejemplo.txt")
	if err != nil {
		fmt.Printf("Error creating dile: %v\n", err)
		return
	}
	defer file.Close()
}
