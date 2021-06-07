package main

import "fmt"

func main() {

	// Long method
	// i := 1
	// for i <= 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// Short method
	for i := 1; i <= 10; i++{
		fmt.Printf("Numero %d\n", i)
	}

	// FizzBuzz: imprime Fizz si es multiplo de 3, Buzz si de 5, y FizzBuzz si es múltiplo de 3 y de 5

	for i := 1; i <= 100; i++ {

		fmt.Printf("%d : ", i)

		if i % 15  == 0 { // Si es divisible por 15
			fmt.Println("Es múltiplo de 15")

		}else if i % 5 == 0 {
			fmt.Println("Es múltiplo de 5")

		}else if i % 3 == 0{
			fmt.Println("Es múltiplo de 3")

		}else{
			fmt.Println("No es múltiplo ni de 3 ni de 5")

		}
	}
}
