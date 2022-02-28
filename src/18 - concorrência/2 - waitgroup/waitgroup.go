package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() { //"go" executa de forma concorrente as duas funções abaixo

		escrever("Olá Mundo!")
		waitGroup.Done() // -1 em waitGroup.Add

	}()

	go func() {

		escrever("Programando em Go")
		waitGroup.Done() // -1 em waitGroup.Add

	}()

	waitGroup.Wait() // esperar waiiGroup.Add chegar em 0

}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
