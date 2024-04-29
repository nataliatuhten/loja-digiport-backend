package main

import "fmt"

func main() {
	//catalogoprodutos := ListaProdutos()
	//fmt.Printf("Este é o catálogo da loja GlowGlam Cosmetics:, %+v\n", catalogoprodutos)
	//produtosCatalogo := ListaProdutos()
	//fmt.Printf("Esse é o catalogo da loja: %+v\n\n", produtosCatalogo)

	fmt.Println("Esse é o catálogo da loja GlowGlam Cosmetics: ")

	nome := "Batom BT Angel" // substituir por um nome de produto

	produtosFiltrados := buscaProdutosPorNome(nome)

	fmt.Println(produtosFiltrados)
	StartServer()
}
