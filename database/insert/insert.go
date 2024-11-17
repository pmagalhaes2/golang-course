package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgres://postgres:postgres@localhost:5432/golang_course?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Erro ao conectar no banco de dados:", err)
	}
	defer db.Close()

	stmt, _ := db.Prepare("INSERT INTO usuarios(nome) VALUES($1)")
	stmt.Exec("Maria")
	stmt.Exec("João")

	res, _ := stmt.Exec("Bruna")
	id, _ := res.LastInsertId()
	fmt.Println(id)

	linhas, _ := res.RowsAffected()
	fmt.Printf("%d linhas afetadas!\n", linhas)
}
