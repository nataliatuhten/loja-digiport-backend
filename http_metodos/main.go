package main

import "net/http"

func main() {
	rotas := Rotas()
	http.ListenAndServe(":8080", rotas)
}
