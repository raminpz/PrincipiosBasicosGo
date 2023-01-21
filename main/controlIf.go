package main

import "fmt"

func main() {
	numero := 0
	if numero == 40 {
		fmt.Println("Los numeros no son iguales")
	} else if numero == 50 {
		fmt.Println("Los numeros son iguales")
	} else {
		fmt.Println("Como final todos son falsos")
	}

	edad := 18

	if edad >= 18 {
		fmt.Println("Eres mayor de edad")
	} else if edad < 18 && edad > 0 {
		fmt.Println("Eres menor de edad")
	} else {
		fmt.Println("La edad no es válida")
	}
}
