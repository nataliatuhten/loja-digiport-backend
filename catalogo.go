package main

import "github.com/nataliatuhten/loja-digiport-backend/model"

func ListaProdutos() []model.Produto {
	produtos := []model.Produto{
		{
			ID:         "5T9hR7",
			Nome:       "Base Revlon",
			Descricao:  "Base líquida",
			Categoria:  "Base",
			Valor:      89.90,
			Quantidade: 6,
			Imagem:     "base.png",
		},
		{ID: "K2p3QF",
			Nome:       "Batom BT Angel",
			Descricao:  "Batom matte",
			Categoria:  "Batom",
			Valor:      34.90,
			Quantidade: 12,
			Imagem:     "batom.png",
		},
		{ID: "b8Ls4E",
			Nome:       "Paleta de sombra Mari Saad",
			Descricao:  "Paleta de sombras",
			Categoria:  "Sombra",
			Valor:      59.90,
			Quantidade: 4,
			Imagem:     "paleta.png"},

		{ID: "6jFp9A",
			Nome:       "Corretivo BT",
			Descricao:  "Corretivo Líquido",
			Categoria:  "Corretivo",
			Valor:      46.90,
			Quantidade: 6,
			Imagem:     "corretivo.png"},

		{ID: "D4cR8W",
			Nome:       "Máscara de cílios Maybelline",
			Descricao:  "Máscara de cílios",
			Categoria:  "Mascara de Cílios",
			Valor:      39.90,
			Quantidade: 12,
			Imagem:     "rimel.png"},
	}
	return produtos
}
