package main

import "fmt"



func main() {
	// Arrays
	var fruitArray [2] string

	// Assign values
	fruitArray[0] = "Apple"
	fruitArray[1] = "Orange"
	
	fmt.Println(fruitArray) // Output: [Apple Orange]

	// Declare and assign
	fruitArray2 := [2]string{"Manzana", "Naranja"}
	fmt.Println(fruitArray2)

	// Slices
	fruitSlice := []string{"Plátano", "Pera", "Frambuesa"}
	fmt.Println(fruitSlice)

	// Tamaño de un array/slice
	fmt.Println("Tamaño del slice: ",len(fruitSlice))

	// Intervalos
	fmt.Println(fruitSlice[1:3]) //Output: [Pera Frambuesa]
}
