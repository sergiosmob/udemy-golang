package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Página raiz")) // similar ao "Println", não aceito aqui no http
}

func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carregar página de usuários")) // similar ao "Println", não aceito aqui no http
}

func main() {

	// HTTP: protocolo
	// Request(cliente) - Response(servidor)
	// Rotas (pra onde ai a requisição? Cadastro, compra de produto...etc)
	// URI - Identificador do recurso
	// Método: GET, POST, PUT, DELETE

	// 1 - Declarando as rotas:

	http.HandleFunc("/home", home)
	http.HandleFunc("/usuarios", usuarios)

	// 2 - subindo um servidor http em Go na porta 5000:

	log.Fatal(http.ListenAndServe(":5000", nil))

}
