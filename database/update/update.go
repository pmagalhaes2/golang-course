package main

import (
	"database/sql"
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

	// UPDATE
	stmt, _ := db.Prepare("UPDATE usuarios SET nome = $1 WHERE id = $2")
	stmt.Exec("Bernardo", 3)

	// DELETE
	stmt2, _ := db.Prepare("DELETE FROM usuarios WHERE id = $1")
	stmt2.Exec(4)
}
