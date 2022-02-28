package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{ // dentro do colchetes define-se o tipo das chaves. Fora dele o tipo de valores
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"ultimo":   "Pedro",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "Campus 1",
		},
	}
	fmt.Println(usuario2)
	//delete(usuario2, "nome")
	//fmt.Println(usuario2)
	usuario2["Signo"] = map[string]string{
		"nome": "Gêmeos",
	}
	fmt.Println(usuario2)
}
