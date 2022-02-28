package main

import (
	"errors"
	"fmt"
)

func main() {
	numero := 100000000000
	numero1 := -10000000
	fmt.Println(numero, numero1)

	var numero2 uint32 = 10000 //"uint não aceita sinal = unusign"
	fmt.Println(numero2)

	//alias
	//int32 = rune

	var numero3 rune = 123456
	fmt.Println(numero3)

	//alias
	//int8 = byte (1 byte - 8 bites)

	var numero4 byte = 123
	fmt.Println(numero4)

	// Float (somente float32 e float64)

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123000000.45
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)

	// fim de exemplos núemros reais

	var str string = "juajuajuajuajuajuhuahuhaua"
	fmt.Println(str)

	str2 := "texto 2"
	fmt.Println(str2)

	char := 'B'
	//char com áspas simples vai trazer o nº do caractere na tabela ASSCII,
	// que no caso é 66 - ele é do tipo "int"
	fmt.Println(char)

	//FIM STRINGS

	var texto int16
	fmt.Println(texto) // var retornar o valor "0". No Go todo tipo de dado tem o valor inicial "0"
	//no caso string, o valor retornado vai ser vazio.

	var boleano1 bool = true //se deixar sem, o valor "0" do boleano é "false"
	fmt.Println(boleano1)

	var erro1 error = errors.New("Erro interno") // o valor "0" do tipo é "nil"
	// erro1 = variável
	// error = tipo
	// errors = pacote nativo do go
	fmt.Println(erro1)

}
