package main

import "fmt"

func greeting(name string) string{
	return "Hola " + name + "!"
}

func getSum(num1 int, num2 int) int{
	return num1 + num2
}

func main() {
	fmt.Println(greeting("Javi"))
	fmt.Println(getSum(10, 12))
}
