package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func InitDB() *sql.DB {
	// Carrega o tipo do banco
	db, err := sql.Open("sqlite", "./pedidos.db")
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}

	// Verifica conexão com o banco
	if err := db.Ping(); err != nil {
		log.Fatalf("erro connection to SQLlite: %v", err)
	}

	// Criação da tabela se não existe
	query := `
	CREATE TABLE IF NOT EXISTS pedidos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		description TEXT NOT NULL,
		status TEXT NOT NULL DEFAULT 'aberto',
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err = db.Exec(query)
	if err != nil {
		log.Fatalf("error on creating table: %v", err)
	}

	return db
}
