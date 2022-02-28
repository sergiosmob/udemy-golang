package main

import (
	"fmt"
	"time"
)

func main() {

	//i := 0
	//for i < 10 {
	//	i++
	//	fmt.Println("Incrementando i")
	//	time.Sleep(time.Second)
	//}
	//fmt.Println(i)

	//= 0; j < 10; j ++ { // inclrementa de 1 em 1
	//0; j < 10; j += 2 {
	//rintln("Incrementando j", j)
	//Sleep(time.Second)
	//	}

	nomes := [3]string{"João", "David", "Lucas"} // o primeiro parâmetro é sempre o índice (0, 1, 2, 3 - no caso)

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for _, nome := range nomes { // vai retornar apenas o nome colocando o "underline no lugar do índice e removento o parâmetro"
		fmt.Println(nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{ // Não é possível usar range em "structs"
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}
	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	// loop infinito

	for {
		fmt.Println("Executando infinitamente")
		time.Sleep(time.Second)
	}
}
