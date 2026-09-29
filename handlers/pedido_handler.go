package handlers

import (
	"database/sql"
	"fmt"
	"html/template"
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

func (h PedidoHandler) List(w http.ResponseWriter, r *http.Request) {
	pedidos, err := models.ListOrders(h.DB)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if r.Method != http.MethodGet {
		fmt.Println(r)
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("GET")

	tmpl := template.Must(template.ParseFiles("template/index.html"))
	tmpl.Execute(w, pedidos)
}
