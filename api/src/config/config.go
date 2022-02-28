package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (

	// StringConexaoBando é a string de conexão com o MySql
	StringConexaoBanco = ""

	// Porta onde a API vai rodar
	Porta = 0

	// SecretKey é a chave que vai assinar o token
	SecretKey []byte
)

//Carregar vai inicializar as variáveis de ambiente
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro) // mata a execução se der problema
	}

	// pega a string e converte para inteiro
	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Porta = 9000
	}

	StringConexaoBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NOME"),
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
