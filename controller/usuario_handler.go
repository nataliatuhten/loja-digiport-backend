package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nataliatuhten/loja-digiport-backend/model"
)

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
