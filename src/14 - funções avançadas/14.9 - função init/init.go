package main

import "fmt"

func main() {
	fmt.Println("Função main sendo executada")
}

//a func init vai ser executada antes da func main, mesmo fora de ordem.
//função utilizada para configurar algum setup por ex.
func init() {

	fmt.Print("Executando a função init")
}
