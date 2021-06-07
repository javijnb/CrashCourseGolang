package main

import ("fmt"
		"io/ioutil"
		"log"
		"net/http"
)

func main() {

	// fmt.Println("1 - INICIO")
	// defer fmt.Println("2 - MEDIO")
	// fmt.Println("3 - FIN")

	res, err := http.Get("http://www.google.com/robots.txt")

	if err != nil {
		log.Fatal(err)
	}

	// Permite indicar cuanto antes que se cierre el lector, pero no que se cierre en este momento,
	// sino cuando acabe la ejecución correcta del programa
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Robots: %s", robots)

}
