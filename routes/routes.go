package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nataliatuhten/loja-digiport-backend/controller"
)

func HandleRequests() {
	route := mux.NewRouter()
	route.HandleFunc("/produtos", controller.BuscaProdutosHandler).Methods("GET")
	route.HandleFunc("/produto", controller.BuscaProdutoPorNomeHandler).Methods("GET")
	route.HandleFunc("/produtos", controller.CriaProdutosHandler).Methods("POST")
	http.ListenAndServe(":8080", route)
}
