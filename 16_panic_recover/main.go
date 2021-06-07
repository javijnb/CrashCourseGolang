package main

import "fmt"
import "log"

func main() {
	fmt.Println("Inicio main")
	panico()
	fmt.Println("Fin main")
}

func panico() {
	fmt.Println("A punto de entrar en pánico")

	

	//Esto se pospone (hasta fin de ejecución o un panic)
	defer fmt.Println("Print defer")
	defer func() {

		if err := recover(); err != nil {
			log.Println("Error: ", err)
			//Si necesitamos propagar el error ya a la rutina de ejecución, lanzamos aquí un panic, fuera del recover
			// panic(err)
		}
	}() // <--- Así forzamos a que se ejecute la función anónima

	

	panic("Algo malo ha pasado...")
	fmt.Println("Fin del pánico") //Unreachable code
}
