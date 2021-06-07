package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a,b)
	fmt.Printf("%T\n", b) // Output: *int

	// Usar *b para leer los valores apuntados por una variable puntero
	fmt.Println(*b)
	fmt.Println(*&a)

	// Cambiar un valor apuntado
	*b = 10
	fmt.Println(a)
}
