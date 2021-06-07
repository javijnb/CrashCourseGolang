package main

import ("fmt"
		"math")

// Interfaz:
type Forma  interface {
	area() float64
}

// Formas a usar:
type Circulo struct{
	x, y, radius float64
}

type Rectangulo struct{
	ancho, alto float64
}

// Métodos para calcular area:
func (c Circulo) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangulo) area() float64{
	return r.ancho * r.alto
}

// Conseguir area:
func getArea(f Forma) float64{
	return f.area()
}

func main() {
	circulo1 := Circulo{ x:0, y:0, radius:5 }
	rectangulo1 := Rectangulo{ ancho:10, alto:5 }

	fmt.Printf("Área del círculo: %f\n", getArea(circulo1))
	fmt.Printf("Área del rectánuglo: %f\n", getArea(rectangulo1))
}
