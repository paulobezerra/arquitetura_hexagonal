package db

import (
	"database/sql"
	"log"

	"github.com/bezerrapaulo/go-hexagonal/application"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDB struct {
	db *sql.DB
}

func NewProductDB(db *sql.DB) *ProductDB {
	return &ProductDB{db: db}
}

func (p *ProductDB) Get(ID string) (application.ProductInterface, error) {
	var product application.Product

	stmt, err := p.db.Prepare("select id, name, status, price from products where id=?")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	err = stmt.QueryRow(ID).Scan(&product.ID, &product.Name, &product.Status, &product.Price)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *ProductDB) Save(product application.ProductInterface) (application.ProductInterface, error) {

	var rows int
	err := p.db.QueryRow("select count(*) as rows from products where id=?", product.GetID()).Scan(&rows)

	if err != nil {
		log.Fatal(err.Error())
		return &application.Product{}, err
	}

	if rows == 0 {
		return p.create(product)
	} else {
		return p.update(product)
	}
}

func (p *ProductDB) create(product application.ProductInterface) (application.ProductInterface, error) {
	stmt, err := p.db.Prepare(`insert into products(id, name, status, price) values(?, ?, ?, ?)`)

	if err != nil {
		return &application.Product{}, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(product.GetID(), product.GetName(), product.GetStatus(), product.GetPrice())

	if err != nil {
		return &application.Product{}, err
	}

	return product, nil
}

func (p *ProductDB) update(product application.ProductInterface) (application.ProductInterface, error) {
	stmt, err := p.db.Prepare(`update products set name=?, status=?, price=? where id=?`)

	if err != nil {
		return &application.Product{}, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(product.GetName(), product.GetStatus(), product.GetPrice(), product.GetID())

	if err != nil {
		return &application.Product{}, err
	}

	return product, nil
}
