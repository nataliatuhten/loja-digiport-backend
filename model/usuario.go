package model

import (
	"fmt"

	"github.com/nataliatuhten/loja-digiport-backend/db"
	"golang.org/x/crypto/bcrypt"
)

type Usuario struct {
	ID       string `json:"id"`
	Nome     string `json:"nome"`
	Email    string `json:"email"`
	Senha    string `json:"senha"`
	Telefone string `json:"telefone"`
	Endereço string `json:"endereco"`
}

func CriaUsuario(usuario Usuario) error {
	hash, err := hashPassword(usuario.Senha)
	if err != nil {
		return fmt.Errorf("erro ao criar o usuário: %w", err)
	}

	db := db.ConectaBancoDados()
	_, err = db.Exec("INSERT INTO USUARIO(nome, senha, email, telefone, endereco) VALUES ($1, $2, $3, $4, $5)",
		usuario.Nome, hash, usuario.Email, usuario.Telefone, usuario.Endereço)
	if err != nil {
		return fmt.Errorf("erro ao inserir usuário no banco de dados: %w", err)
	}
	return nil
}

func hashPassword(senha string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(senha), 14)
	if err != nil {
		return "", fmt.Errorf("erro ao tentar gerar hash da senha: %w", err)
	}
	return string(bytes), nil
}
