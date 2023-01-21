package main

import "fmt"

func main() {

	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	nombre := "Ramiro"
	fmt.Println(string(nombre[5]))

	for i := 0; i < 7; i++ {
		fmt.Println(i, string(nombre[i]))
	}

	// En GO no existe while
	numero := 1
	for numero != 1 {
		fmt.Println("Ejecuto el código dentro del FOR")
	}
	fmt.Println("Texto despúes del FOR")

}
