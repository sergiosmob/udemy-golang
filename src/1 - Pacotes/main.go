package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do pacote main")
	auxiliar.Escrever()
	erro := checkmail.ValidateFormat("123")
	fmt.Println(erro)
}
