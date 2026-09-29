package handlers

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"pedidoservice/models"
	"strconv"
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

func (h PedidoHandler) Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost && r.Method != http.MethodGet {
		fmt.Println(r)
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	stringId := r.URL.Query().Get("id")
	fmt.Println(stringId)

	id, err := strconv.Atoi(stringId)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
	}

	err = models.EditOrder(h.DB, id)
	if err != nil {
		http.Error(w, "error updating order: "+err.Error(), http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h PedidoHandler) Remove(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost && r.Method != http.MethodGet {
		fmt.Println(r)
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	stringId := r.URL.Query().Get("id")
	fmt.Println(stringId)

	id, err := strconv.Atoi(stringId)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
	}

	err = models.RemoveOrder(h.DB, id)
	if err != nil {
		http.Error(w, "error updating order: "+err.Error(), http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
