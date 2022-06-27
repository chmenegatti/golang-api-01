package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/chmenegatti/golang-api-01/models"
)

func ListAll(w http.ResponseWriter, r *http.Request) {
	todos, err := models.GetAll()
	if err != nil {
		log.Printf("Erro ao obter registros: %v", err)
	}

	w.Header().Add("Content-Type", "application/josn")
	json.NewEncoder(w).Encode(todos)
}
