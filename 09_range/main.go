package main

import "fmt"

func main() {
	IDs := []int{12, 34, 56, 78, 90, 13, 4}

	// Iterar el array con RANGE
	fmt.Println("-- Usando RANGE --")
	for i, id := range IDs {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Sin usar el index
	fmt.Println("-- Sin usar index --")
	for _, id := range IDs {
		fmt.Printf("ID: %d\n", id)
	}

	// Añadir todos los ID juntos
	fmt.Println("-- Suma de IDs --")
	sum := 0
	for _, id := range IDs{
		sum += id
	}
	fmt.Printf("Suma de IDs: %d\n", sum)

	// Range con mapas
	telefonos := map[string]string{"Bob":"123456789", "Javi":"132457698", "Mike":"192837465"}
	
	for k, v := range telefonos {
		fmt.Printf("Teléfonos: %s - %s\n", k, v)
	}
}
