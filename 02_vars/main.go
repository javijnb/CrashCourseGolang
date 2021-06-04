package main

import "fmt"

var nombreGlobal string = "Javier Núñez Bon"

func main() {

	// Using var: ambos son válidos, incluso sin tipar
	var nombre string = "Javier Núñez"
	var nombre2 = "Javi Núñez - nombre 2"

	var edad int = 22

	fmt.Println(nombre, edad)
	fmt.Println(nombre2)

	// Para ver en qué formato se ha creado una variable:
	fmt.Printf("%T\n", nombre)
	fmt.Println("-------------------------------")

	// Using const: su valor no puede ser modificado
	var mola = true
	mola = false //(Se puede alterar sobre la marcha)

	const noMola = false
	// noMola = true (MAL)
	fmt.Println(nombre, edad, mola, noMola)
	fmt.Println("-------------------------------")

	// Using global vars: siempre var, func, etc), no valen shorthands:
	// Shorthand:
	aux := 10
	fmt.Println("Nombre variable global: ", nombreGlobal)
	fmt.Println("Variable auxiliar: ", aux)

	// Atajos de declaraciones:

	residencia, email := "Pontevedra", "javijnb@gmail.com"
	fmt.Println("Lugar de residencia: ", residencia, ". Correo electrónico: ", email)
	fmt.Println("-------------------------------")

}
