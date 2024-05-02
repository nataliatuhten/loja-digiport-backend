package main

import (
	"encoding/json"
	"net/http"

	"github.com/nataliatuhten/loja-digiport-backend/model"
)

/*
	func StartServer() {
		http.HandleFunc("/produtos", produtosHandler)
		http.ListenAndServe(":8080", nil)
	}

	func produtosHandler(w http.ResponseWriter, r *http.Request) {
		produtos := ListaProdutos()
		json.NewEncoder(w).Encode(produtos)
	}
*/
func StartServer() {
	http.HandleFunc("/produtos", produtosHandler)

	http.ListenAndServe(":8080", nil)
}

func produtosHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		getProdutos(w, r)
	} else if r.Method == "POST" {
		addProduto(w, r)
	}
}

func addProduto(w http.ResponseWriter, r *http.Request) {
	var produto model.Produto
	json.NewDecoder(r.Body).Decode(&produto)

	criarCatalogo(produto)
	err := criarCatalogo(produto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.Erro{ErrorMessage: err.Error()})
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func getProdutos(w http.ResponseWriter, r *http.Request) {
	queryNome := r.URL.Query().Get("nome")
	if queryNome != "" {
		produtosFitradosPorNome := buscaProdutosPorNome(queryNome)
		json.NewEncoder(w).Encode(produtosFitradosPorNome)
	} else {
		produtos := Produtos
		json.NewEncoder(w).Encode(produtos)
	}
}
