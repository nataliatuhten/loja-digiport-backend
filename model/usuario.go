package model

import (
	"database/sql"
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

func ValidaUsuario(usuario Usuario) error {
	if usuario.Email == "" {
		return fmt.Errorf("email do usuário não pode ser vazio")
	}
	if usuario.Senha == "" {
		return fmt.Errorf("senha do usuário não pode ser vazia")
	}
	if usuario.Nome == "" {
		return fmt.Errorf("nome do usuário não pode ser vazio")
	}
	return nil
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

func BuscaUsuarioPorEmail(email string) (*Usuario, error) {
	db := db.ConectaBancoDados()
	defer db.Close()

	var usuario Usuario
	err := db.QueryRow("SELECT id,  nome, senha, email, telefone, endereco FROM usuario WHERE email = $1", email).Scan(&usuario.ID,
		&usuario.Nome, &usuario.Senha, &usuario.Email, &usuario.Telefone, &usuario.Endereço)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("usuario não encontrado %s", email)
		}
		return nil, err
	}
	return &usuario, nil
}

func UpdateUsuario(usuario Usuario) error {
	db := db.ConectaBancoDados()
	defer db.Close()

	id := usuario.ID
	nome := usuario.Nome
	senha := usuario.Senha
	email := usuario.Email
	telefone := usuario.Telefone
	endereco := usuario.Endereço

	result, err := db.Exec("UPDATE usuarios SET nome= $1, senha= $2, email = $3, telefone= $4, endereco=$5 where id= $6", nome, senha, email, telefone, endereco, id)
	if err != nil {
		panic(err.Error())
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}

	if rowsAffected == 0 {
		return fmt.Errorf("produto não encontrado")
	}

	fmt.Printf("Produto %s atualizado com sucesso (%d row affected)\n", id, rowsAffected)

	return nil
}

func ValidaLogin(hash string, senhatxt string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(senhatxt))
	if err != nil {
		return fmt.Errorf("Senha incorreta")
	}
	return nil
}
