package main

import "fmt"

func Exercicio2() {
	minhasDespesasFixas := []string{"aluguel", "internet", "celular"}
	custosDespesasFixas := []float64{}
	custosDespesasFixas = append(custosDespesasFixas, 1500.22, 109.90, 60)

	var total float64
	for i, nomeDespesa := range minhasDespesasFixas {
		total += custosDespesasFixas[i]
		fmt.Println(nomeDespesa, "no valor de", custosDespesasFixas[i])
	}
	fmt.Println("Custo total das despesas fixas:", total)
	orcamento := 1700.00
	fmt.Println("Estamos dentro do orçamento (R$", orcamento, ")")
	if total <= orcamento {
		fmt.Println("✅")
	} else {
		fmt.Println("❌")
	}
}
