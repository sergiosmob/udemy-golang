package main

import "fmt"

func main() {
	canal := make(chan string, 2) //mais de dois valores ocorrer[a deadlock!]
	canal <- "Ola Mundo\n"
	canal <- "Programendo em Go\n"
	mensagem := <-canal
	mensagem2 := <-canal
	fmt.Print(mensagem)
	fmt.Print(mensagem2)
}
