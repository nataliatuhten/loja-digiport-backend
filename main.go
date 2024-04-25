package main

import "fmt"

func main() {
	//catalogoprodutos := ListaProdutos()
	//fmt.Printf("Este é o catálogo da loja GlowGlam Cosmetics:, %+v\n", catalogoprodutos)
	produtosCatalogo := ListaProdutos()
	//fmt.Printf("Esse é o catalogo da loja: %+v\n\n", produtosCatalogo)

	fmt.Println("Esse é o catálogo da loja GlowGlam Cosmetics: ")

	for _, produtos := range produtosCatalogo {
		fmt.Printf("ID: %s\n", produtos.ID)
		fmt.Printf("Nome: %s\n", produtos.Nome)
		fmt.Printf("Descrição: %s\n", produtos.Descricao)
		fmt.Printf("Categoria: %s\n", produtos.Categoria)
		fmt.Printf("Valor: %.2f\n", produtos.Valor)
		fmt.Printf("Quantidade: %d\n", produtos.Quantidade)
		fmt.Printf("Imagem: %s\n", produtos.Imagem)
		fmt.Println("-------------------------------------------")
	}

	nome := "Batom BT Angel" // substituir por um nome de produto

	produtosFiltrados := produtosPorNome(nome)

	fmt.Println(produtosFiltrados)
	StartServer()
}
