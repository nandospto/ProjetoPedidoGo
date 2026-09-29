package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"pedidoservice/models"
)

type PedidoHandler struct {
	DB *sql.DB
}

func (h PedidoHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Println(r)
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("POST")
	descricao := r.FormValue("descricao")
	fmt.Println(descricao)
	if descricao == "" {
		http.Error(w, "Descrição obrigatória", http.StatusBadRequest)
		return
	}

	if err := models.CreateOrder(h.DB, descricao); err != nil {
		http.Error(w, "Erro ao criar pedido", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
