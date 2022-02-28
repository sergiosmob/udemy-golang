package main

import "fmt"

func diaDaSemana(numero int) string {

	switch numero {
	case 1:
		return "domingo"
	case 2:
		return "segunda-feira"
	case 3:
		return "terça -feira"
	case 4:
		return "quarta-feira"
	case 5:
		return "quinta-feira"
	case 6:
		return "sexta-feira"
	case 7:
		return "sabado"
	default:
		return "Número invalido"
	}
}

func return2(numero int) string {
	var diaDaSemana string

	switch {
	case numero == 1:
		diaDaSemana = "domingo"
		fallthrough // quando definida ela vai jogar o valor da variável "diaDaSemana = "domingo" para "segunda-feira" (no caso o próximo) "
	case numero == 2:
		diaDaSemana = "segunda-feira"
	case numero == 3:
		diaDaSemana = "terça-feira"
	case numero == 4:
		diaDaSemana = "quarta-feira"
	case numero == 5:
		diaDaSemana = "quinta-feira"
	case numero == 6:
		diaDaSemana = "sexta-feira"
	case numero == 7:
		diaDaSemana = "sabado"
	}

	return diaDaSemana
}

func main() {
	fmt.Println("switch")
	dia := diaDaSemana(200)
	fmt.Println(dia)

	dia2 := return2(1)
	fmt.Println(dia2)
}
