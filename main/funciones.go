package main

import "fmt"

func main() {
	mi_func()
	mi_func()
	mi_func()
	mi_func()

	saludar(" Ramiro")

	sumar(5, 25)
	fmt.Println(sumar(100, 200))

	suma := sumar(10, 20)
	fmt.Println(suma)

}

func mi_func() {
	fmt.Println("Hola Mundo saludos")
	fmt.Println("Hola Mundo saludos")
	fmt.Println("Hola Mundo saludos")
}

func saludar(nombre string) {
	fmt.Println("Hola" + nombre)
}

func sumar(num1 int, num2 int) int {
	suma := num1 + num2
	return suma

}
