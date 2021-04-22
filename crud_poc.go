package main

import (
	"database/sql" //https://golang.org/pkg/database/sql/
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3" //https://stackoverflow.com/questions/21220077/what-does-an-underscore-in-front-of-an-import-statement-mean
)

//http://go-database-sql.org/overview.html

//Product struct
type products struct {
	id      int
	product string
	ptype   string
	price   int
	os      string
}

//checking error function
func errchecker(err error) {
	if err != nil {
		panic(err)
	}
}

//CREATE and add products in db
func addProduct(db *sql.DB, product string, ptype string, price int, os string) {
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into products (product,ptype,price,os) values (?,?,?,?)")
	_, err := stmt.Exec(product, ptype, price, os)
	errchecker(err)
	tx.Commit()
}

//READ and get data from db
func getProducts(db *sql.DB, id int) products {
	rows, err := db.Query("select * from products")
	errchecker(err)
	for rows.Next() {
		var tempProduct products
		err = rows.Scan(&tempProduct.id, &tempProduct.product, &tempProduct.ptype, &tempProduct.price, &tempProduct.os)
		errchecker(err)
		if tempProduct.id == id {
			return tempProduct
		}
	}
	return products{}
}

//UPDATE data in db
func updateProduct(db *sql.DB, id2 int, product string, ptype string, price int, os string) {
	sprice := strconv.Itoa(price) // int to string
	sid := strconv.Itoa(id2)      // int to string
	tx, _ := db.Begin()

	stmt, _ := tx.Prepare("update products set product=?,ptype=?,price=?,os=? where id=?")
	_, err := stmt.Exec(product, ptype, sprice, os, sid)
	errchecker(err)
	tx.Commit()
}

//Delete product from db
func delProduct(db *sql.DB, id int) {
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("delete from products where id=?")
	_, err := stmt.Exec(id)
	errchecker(err)
	tx.Commit()

}

func main() {
	db, err := sql.Open("sqlite3", "database/qh.db")

	fmt.Println(err)

	//addProduct(db, "WebSecurity", "Antivirus", 500, "Windows")

	//updateProduct(db, 7, "ServerSecurity", "Server", 300, "Unix")

	//delProduct(db, 8)

	for i := 1; i <= 20; i++ {
		fmt.Println(getProducts(db, i))
	}

	/*
		var (
			id      int
			product string
		)
		rows, err := db.Query("select id, product from products where id = ?", 1)
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&id, &product)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(id, product)
		}*/

}
