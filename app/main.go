package main

import (
	"fmt"
	"net/http"
	"pedidoservice/db"
	"pedidoservice/handlers"
)

type Dados struct {
	Mensagem string `json:"mensagem"`
}

func main() {
	// Inicia o Banco de dados
	db := db.InitDB()
	defer db.Close()

	// Inicializa o Handler
	pedidoHandler := &handlers.PedidoHandler{DB: db}

	// Inicializa o server com o template
	http.Handle("/", http.FileServer(http.Dir("template")))

	// Rota Post
	http.HandleFunc("/create", pedidoHandler.Create)

	fmt.Println("Conectado")

	// Porta de conexão com a API
	http.ListenAndServe(":5500", nil)
}
