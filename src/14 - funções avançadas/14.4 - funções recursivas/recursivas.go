package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	fmt.Println("Funções recursivas")
	//1 1 2 3 5 8 13

	posicao := uint(10)
	fmt.Println(fibonacci(posicao))

	// ex de estouro de pilha abaixo
	for i := uint(1); 1 <= posicao; i++ {
		fmt.Println(fibonacci(i))
	}
}

// não é recomendada devido a possívl estouro de pilha (stack overflow)
