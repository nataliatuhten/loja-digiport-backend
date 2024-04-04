package model

type Carrinho struct {
	ID           string
	UserID       string
	ProdutoID    []string
	Quantidade   int
	ValorTotal   float64
	InfoProdutos []InfoProdutos
}

type InfoProdutos struct {
	ProdutoID           string
	QuantidadeDeProduto int
}
