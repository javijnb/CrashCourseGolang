package main

import "fmt"

func main() {
	x := 5
	y := 10

	// IF ELSE STATEMENT
	if x <= y {
		fmt.Printf("%d es menor o igual que %d\n", x, y)
	} else {
		fmt.Printf("%d es menor que %d\n", y, x)
	}

	// ELSE IF STATEMENT
	color := "verde"

	if color == "rojo" {
		fmt.Println("El color es rojo")
	}else if color == "azul"{
		fmt.Println("El color es azul")
	}else{
		fmt.Println("El color no es ni rojo ni azul")
	}

	// SWITCH STATEMENT
	switch color {
		case "rojo":
			fmt.Println("SW: ROJO")
		case "azul":
			fmt.Println("SW: AZUL")
		default:
			fmt.Println("SW: ni rojo ni azul")
	}

}
