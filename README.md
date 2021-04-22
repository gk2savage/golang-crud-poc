# golang-crud-poc #

[![GoDoc](https://img.shields.io/static/v1?label=godoc&message=reference&color=blue)](https://golang.org/pkg/database/sql/)

### Proof of Concept
Implementation of CRUD operations with Go Language using sqlite3 db

- We use a sqlite3 database to perform CRUD operations
- Use of various CRUD functions defined and called in main()
- Use of transactions and Prepared Statements

### Implementation?

go get github.com/mattn/go-sqlite3

The special thing about this library is that it uses the internal sql package of Go. We usually import ''database/sql'' and use sql to execute database queries on the database.

### Source of References:

Package sql provides a generic interface around SQL (or SQL-like) databases. 
https://golang.org/pkg/database/sql/

Overview and Understanding
http://go-database-sql.org/overview.html

Query vs Exec vs Prepare
https://medium.com/@alok.sinha.nov/query-vs-exec-vs-prepare-in-golang-e7c49212c36c

-Girish
