package main

import "fmt"

func closure() func() { //funções que referenciam variáveis que estão fora do seu corpo.
	texto := "Dentro da função closure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {

	texto := "Dentro da função main"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()
}
