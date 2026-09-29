package models

import (
	"database/sql"
	"fmt"
	"time"
)

func CreateOrder(db *sql.DB, descricao string) error {
	_, err := db.Exec(
		`INSERT INTO pedidos
		(description, status, created_at)
		VALUES
		($1, 'aberto', $2)`,
		descricao, time.Now())
	fmt.Println("created")
	return err
}
