package main

import "fmt"

func main() {
	// Nuevo mapa[key]value
	emails := make(map[string]string)

	// Asignamos clave - valores
	emails["Bob"]  = "bob@gmail.com"
	emails["Javi"] = "javijnb@gmail.com" 
	emails["Mike"] = "mike@gmail.com" 

	fmt.Println(emails)
	fmt.Printf("Tamaño emails: %d\n", len(emails))
	fmt.Println(emails["Bob"])

	// Eliminar
	fmt.Println("Eliminamos el correo de Bob")
	delete(emails, "Bob")
	fmt.Println(emails)
	fmt.Printf("Tamaño: %d\n", len(emails))

	// Declarar y asignar clave - valor
	telefonos := map[string]string{"Bob":"123456789", "Javi":"132457698", "Mike":"192837465"}
	telefonos["Sharon"] = "132475689"

	fmt.Println(telefonos)
	fmt.Printf("Tamaño telefonos: %d\n", len(telefonos))
	fmt.Println(telefonos["Bob"])
}
