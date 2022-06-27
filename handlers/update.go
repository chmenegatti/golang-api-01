package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/chmenegatti/golang-api-01/models"
	"github.com/go-chi/chi/v5"
)

func Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Erro ao fazer o parse do id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var todo models.Todo

	err = json.NewDecoder(r.Body).Decode(&todo)

	if err != nil {
		log.Println("Erro ao fazer o decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := models.Update(int64(id), todo)
	if err != nil {
		log.Println("Erro ao fazer o atualizar resgistro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Println("Error: foram atualizado %d registros", rows)
	}

	resp := map[string]any{
		"Error":   false,
		"Message": fmt.Sprintf("dados atualizados com sucesso!"),
	}

	w.Header().Add("Content-Type", "application/josn")
	json.NewEncoder(w).Encode(resp)
}
