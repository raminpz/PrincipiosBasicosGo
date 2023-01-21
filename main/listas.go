package main

import "fmt"

func main() {
	frutas := []string{"Pera", "Manzana", "Platano", "Naranja"}
	fmt.Println(frutas[1]) // Mostracion elemento en la posicion 1

	// Agregar datos a la lista o array, metodo append
	frutas = append(frutas, "Melocoton", "Melon", "Toronja")

	// Cambianos el valor de la posicion[0]
	frutas[0] = "Sandía"

	// Recorrer el array
	for i := 0; i < 6; i++ {
		fmt.Println(frutas[i])
	}

	// Recorrer el array y busqueda de Melocoton
	for i := 0; i < len(frutas); i++ {
		if frutas[i] == "Melocoton" {
			fmt.Println("Hay una coincidencia en la lista")
		}
	}

}
