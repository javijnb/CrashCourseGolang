package main

import ("fmt"
		"strconv")

// Definir struct persona
type Persona struct {
	nombre string
	apellidos string
	ciudad string
	edad int
}

// Value receivers: método saludo
func (p Persona) saludo() string{
	return "Hola, mi nombre es " + p.nombre + " " + p.apellidos + ", vivo en " +p.ciudad+ " y tengo " +strconv.Itoa(p.edad)+ " años"
}

// Pointer receivers: método cumpleaños
func(p *Persona) cumpleaños() {
	p.edad++
}

// Pointer receiver: cambio de ciudad
func(p *Persona) mudanza(destino string){
	p.ciudad = destino
}


func main() {
	// Inicializamos la struct
	persona1 := Persona{nombre: "Javi", apellidos: "NB", ciudad: "Pontevedra", edad: 22}
	persona2 := Persona{"Jorge", "Pérez", "Vigo", 45}

	fmt.Println(persona1)
	fmt.Println(persona2)

	// //Coger un campo en concreto:
	// fmt.Println(persona1.nombre)
	// persona2.edad++
	// fmt.Println(persona2)

	persona1.cumpleaños()
	persona1.mudanza("Vigo")
	fmt.Println(persona1.saludo())

}
