package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/nataliatuhten/loja-digiport-backend/model"
)

var jwtKey = []byte("secret")

func CriaUsuarioHandler(w http.ResponseWriter, r *http.Request) {
	var usuario model.Usuario
	json.NewDecoder(r.Body).Decode(&usuario)

	error := model.CriaUsuario(usuario)

	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(error.Error())
	} else {
		w.WriteHeader(http.StatusCreated)

	}

}

func BuscaUsuarioPorEmail(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")

	usuario, err := model.BuscaUsuarioPorEmail(email)
	if err != nil {
		fmt.Println("Erro ao buscar usuario:", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(usuario)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var usuario model.Usuario
	json.NewDecoder(r.Body).Decode(&usuario)
	username := usuario.Email
	senhatxt := usuario.Senha
	user, error := model.BuscaUsuarioPorEmail(username)
	if error != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	hash := user.Senha
	error = model.ValidaLogin(hash, senhatxt)

	if error == nil {
		dataExpiracao := time.Now().Add(5 * time.Minute)
		standardToken := &jwt.StandardClaims{
			ExpiresAt: dataExpiracao.Unix(),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, standardToken)
		tokenString, err := token.SignedString(jwtKey)

		if err != nil {
			fmt.Println("Erro ao validar jwt")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.Write([]byte(tokenString))
		return

	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
