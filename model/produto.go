package model

type Produto struct {
	ID         string  `json: "id"`
	Nome       string  `json: "nome"`
	Descricao  string  `json: "cescricao"`
	Categoria  string  `json: "categoria"`
	Valor      float64 `json: "valor"`
	Quantidade int     `json: "quantidade"`
	Imagem     string  `json: "imagem"`
}
