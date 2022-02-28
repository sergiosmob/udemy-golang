package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	stringConexao := "golang1:golang1@/devbook?charset=utf8&parseTime=True&loc=Local"
	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		log.Fatal(erro)
	}

	defer db.Close()

	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}

	fmt.Print("A conexão está aberta!\n")

	linhas, erro := db.Query("SELECT * FROM usuarios")
	if erro != nil {
		log.Fatal(erro) // log.Fatal pára a execução de tudo!
	}

	defer linhas.Close()
	fmt.Println(linhas)
}
