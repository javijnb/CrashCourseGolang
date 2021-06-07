package main

import ("fmt"
 		"net/http")

func index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1>Hola Mundo Web</h1>")
}

func home(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1>Página principal del servicio web</h1>")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/home", home)
	fmt.Println("Servidor iniciándose...")
	http.ListenAndServe(":34567", nil)
}
