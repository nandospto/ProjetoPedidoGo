package models

import (
	"database/sql"
	"fmt"
	"time"
)

type Order struct {
	Id          int
	Description string
	Status      string
	Created_at  time.Time
}

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

func ListOrders(db *sql.DB) ([]Order, error) {
	query := `
	SELECT id, description, status, created_at
	FROM pedidos
	ORDER BY id ASC
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pedidos []Order
	for rows.Next() {
		var p Order
		if err := rows.Scan(&p.Id, &p.Description, &p.Status, &p.Created_at); err != nil {
			return nil, err
		}
		pedidos = append(pedidos, p)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return pedidos, nil
}
