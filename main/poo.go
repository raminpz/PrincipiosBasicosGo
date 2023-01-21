package main

import "fmt"

type persona struct {
	nombre    string
	apellidos string
	edad      int
}

// Método
func (p persona) saludar(saludo string) {
	fmt.Println(saludo + ", " + p.nombre)
}

func (pers persona) cumple() int {
	return pers.edad + 1
}

func main() {
	persona1 := persona{"Ramiro", "Nuñez Perez", 35}

	persona2 := persona{"Emily", "Benz", 20}

	fmt.Println("Persona1= ", persona1)
	fmt.Println("Persona2= ", persona2)

	persona2.edad = 36
	fmt.Println("Persona2= ", persona2)

	persona1.saludar("Hello.!")
	persona2.saludar("Hola")

	fmt.Println(persona1.cumple())

	edad_persona2 := persona2.cumple()
	fmt.Println(edad_persona2)
}
