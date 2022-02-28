package main

import "fmt"

func main() {
	fmt.Println("estruturas-de-controle")

	numero := -5
	if numero > 15 {
		fmt.Println("Maior do que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	// if init
	if outroNumero := numero; outroNumero > 0 { // ex. uma variável recebendo (:=) o valor de outra variável e testando se o valor é maior que zero
		fmt.Println("Número é maior que zero")
	} else if numero < -10 {
		fmt.Println("Número é menor que -10")
	} else {
		fmt.Println("Número está entre 0 e -10")
	}
}
