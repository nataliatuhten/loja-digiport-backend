package model

import "github.com/nataliatuhten/loja-digiport-backend/db"

type Produto struct {
	ID                  string  `json: "id"`
	Nome                string  `json: "nome"`
	Preco               float64 `json: "valor"`
	Descricao           string  `json: "cescricao"`
	Imagem              string  `json: "imagem"`
	QuantidadeEmEstoque int     `json: "quantidade"`
}

var id, nome string
var preco float64
var descricao, imagem string
var quantidade int

func BuscaTodosProdutos() []Produto {
	db := db.ConectaBancoDados()
	resultado, err := db.Query("SELECT * FROM produtos")
	err = resultado.Scan(&id, &nome, &preco, &descricao, &imagem, &quantidade)
	if err != nil {
		panic(err.Error())
	}
	p := Produto{}
	produtos := []Produto{}
	for resultado.Next() {
		p.ID = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Imagem = imagem
		p.QuantidadeEmEstoque = quantidade

		produtos = append(produtos, p)

	}

	defer db.Close()
	return produtos
}
