package main

import "fmt"

type Contato struct {
	Nome     string
	Telefone string
	Email    string
}

func main() {
	// Criando alguns contatos
	contato1 := Contato{Nome: "João", Telefone: "123456789", Email: "joao@example.com"}
	contato2 := Contato{Nome: "Maria", Telefone: "987654321", Email: "maria@example.com"}
	// Criando um slice para armazenar os contatos
	contatos := []Contato{contato1, contato2}
	// Iterando sobre os contatos e exibindo suas informações
	for i, contato := range contatos {
		fmt.Printf("Contato %d:\n", i+1) //pois a primeira posição é no indice zero
		fmt.Printf("Nome: %s\nTelefone: %s\nEmail: %s\n\n", contato.Nome, contato.Telefone, contato.Email)
	}
	contato3 := Contato{Nome: "Emily", Telefone: "111111111", Email: "emily@example.com"}
	contatos = append(contatos, contato3)

	for i, contato := range contatos {
		fmt.Printf("Contato %d:\n", i+1)
		fmt.Printf("Nome: %s\nTelefone: %s\nEmail: %s\n\n", contato.Nome, contato.Telefone, contato.Email)
	}
}
