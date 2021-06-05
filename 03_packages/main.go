package main

import (
	"fmt"
	"math"
	"p1"
)

func main() {
	fmt.Println("Estamos usando el paquete MATH: ")
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(16))

	fmt.Println(p1.Reverse("!oít aloH"))
}

func Reverse(s string) string {

	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)

}
