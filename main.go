package main

import (
	"fmt"

	"github.com/nataliatuhten/loja-digiport-backend/model"
)

func main() {
	//catalogoprodutos := ListaProdutos()
	//fmt.Printf("Este é o catálogo da loja GlowGlam Cosmetics:, %+v\n", catalogoprodutos)
	//produtosCatalogo := ListaProdutos()
	//fmt.Printf("Esse é o catalogo da loja: %+v\n\n", produtosCatalogo)

	//--fmt.Println("Esse é o catálogo da loja GlowGlam Cosmetics: ")
	//--nome := "Batom BT Angel" // substituir por um nome de produto
	//--produtosFiltrados := buscaProdutosPorNome(nome)
	//--fmt.Println(produtosFiltrados)

	produtoNovo := model.Produto{
		ID:         "1",
		Nome:       "Batom Mari Saad",
		Descricao:  "Batom Gloss",
		Categoria:  "Batom",
		Valor:      53.99,
		Quantidade: 3,
		Imagem:     "batomarisaad.png",
	}
	err := criarCatalogo(produtoNovo)
	if err != nil {
		fmt.Println("Erro ao adicionar o produto", err)
		return
	}
	//fmt.Println("Produto adicionado com sucesso!")
	StartServer()
}
