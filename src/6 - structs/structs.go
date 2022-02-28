package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco //campo "endereco" do tipo "endereco" - strucrt criado abaixo
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "David"
	u.idade = 21
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua sobe desce", 0}

	usuario2 := usuario{"David", 21, enderecoExemplo}
	fmt.Println(usuario2)
}
