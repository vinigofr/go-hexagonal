package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/vinigofr/go-hexagonal/application"
)

type ProductDb struct {
	db *sql.DB
}

func (p *ProductDb) Get(id string) (application.ProductInterface, error) {
	var product application.Product

	stmt, err := p.db.Prepare("SELECT, id, name, price, status FROM products WHERE id = ?")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(
		&product.ID,
		&product.Price,
		&product.Name,
		&product.Status,
	)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *ProductDb) Save(id string) (application.ProductInterface, error)
