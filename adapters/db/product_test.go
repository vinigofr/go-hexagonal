package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vinigofr/go-hexagonal/adapters/db"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products (
		products(
			"id" string,
			"name" string,
			"price" float,
			"status" string
		);`

	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err)
	}

	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `INSERT INTO products values (
		"abc",
		"test",
		0,
		"disabled"
	);`

	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err)
	}

	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer Db.Close()

	productDb := db.NewProductDb(Db)

	product, err := productDb.Get("abc")
	require.Nil(t, err)

	require.Equal(t, "name", product.GetName())
}
