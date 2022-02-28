package main

import (
	"fmt"
	"time"
)

//CONCORRÊNCIA É DIFERENTE(!=) DE PARALELISMO.
//PARALELISMO = duas ou mais tarefas sendo executadas ao mesmo tempo, distribuidas em vários núcleos do processador
//CONCORRÊNCIA = duas ou mais tarefas sendo executadas, não ao mesmo tempo, num mesmo núcleo, revezendo no processamento

func main() {
	go escrever("Olá Mundo!") // goroutine - não vai esperar o processo terminar para começar o debaixo.
	escrever("Programando em Go")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}

}
