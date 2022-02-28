package main

import "fmt"

func inverterSinal(numero int) int { //passando uma cópia do valor registrado em memória para a função. O seu valor original é preservado.
	return numero * -1
}

func inverterSinalComPonteiro(numero *int) { //passando uma referência para a função. Ou seja, o valor original será alterado.
	*numero = *numero * -1
}

func main() {
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)
	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}
