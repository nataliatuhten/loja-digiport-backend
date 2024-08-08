package controller

import (
	"encoding/json"
	"net/http"

	"github.com/nataliatuhten/loja-digiport-backend/model"
)

func BuscaProdutosHandler(w http.ResponseWriter, r *http.Request) {
	produtos := model.BuscaTodosProdutos()
	json.NewEncoder(w).Encode(produtos)
}

func BuscaProdutoPorNomeHandler(w http.ResponseWriter, r *http.Request) {

}

func CriaProdutosHandler(w http.ResponseWriter, r *http.Request) {
}
