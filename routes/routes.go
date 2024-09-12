package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nataliatuhten/loja-digiport-backend/controller"
	"github.com/rs/cors"
)

func HandleRequests() {
	route := mux.NewRouter()
	route.HandleFunc("/produtos", controller.BuscaProdutosHandler).Methods("GET")
	route.HandleFunc("/produto", controller.BuscaProdutoPorNomeHandler).Methods("GET")
	route.HandleFunc("/produto", controller.CriaProdutosHandler).Methods("POST")
	route.HandleFunc("/produto/{id}", controller.RemoveProdutoHandler).Methods("DELETE")

	//usuario
	route.HandleFunc("/usuarios", controller.CriaUsuarioHandler).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	handler := c.Handler(route)
	http.ListenAndServe(":8080", handler)
}
