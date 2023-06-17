package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/bezerrapaulo/go-hexagonal/adapters/db"
	"github.com/bezerrapaulo/go-hexagonal/application"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp() {
	var err error
	Db, err = sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err.Error())
	}
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products(
			"id" string, 
			"name" string, 
			"price" float, 
			"status" string
			);`

	_, err := db.Exec(table)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func createProduct(db *sql.DB) {
	insert := `INSERT INTO products values("abc", "Product Test", 0, "disabled");`
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer stmt.Close()

	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer Db.Close()
	productDb := db.NewProductDB(Db)

	product, err := productDb.Get("abc")

	require.Nil(t, err)
	require.Equal(t, "abc", product.GetID())
	require.Equal(t, "Product Test", product.GetName())
	require.Equal(t, 0.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}

func TestProductDb_SaveCreate(t *testing.T) {
	setUp()
	defer Db.Close()

	product := application.NewProduct()
	product.Name = "Product Test"
	product.Price = 25.0

	productDb := db.NewProductDB(Db)

	result, err := productDb.Save(product)

	require.Nil(t, err)
	require.Equal(t, product.ID, result.GetID())
	require.Equal(t, product.Name, result.GetName())
	require.Equal(t, product.Price, result.GetPrice())
	require.Equal(t, application.DISABLED, result.GetStatus())

	err = result.Enable()

	require.Nil(t, err)

	result, err = productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.GetID(), result.GetID())
	require.Equal(t, product.GetName(), result.GetName())
	require.Equal(t, product.GetPrice(), result.GetPrice())
	require.Equal(t, application.ENABLED, result.GetStatus())
}
