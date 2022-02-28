package main

import "fmt"

func main() {

	// ARITMETICOS ("+", "-", "*", "/", "\")

	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int16 = 25
	soma2 := numero1 + numero2 // tem que ser do mesmo tipo (no caso ex. int16)
	fmt.Println(soma2)

	//FIM DOS OPERADORES ARITMÉTICOS

	// ATRIBUIÇÃO ("=", ":=")

	var variavel1 string = "String"
	variavel2 := "string2"
	fmt.Println(variavel1, variavel2)

	//FIM DOS OPERADORES DE ATRIBUIÇÃO

	//RELACIONAIS (Retornam sempre boleano)
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2) //comparação
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 > 2)
	fmt.Println(1 != 2)

	//FIM DOS OPERADORES RELACIONAIS

	// LÓGICOS
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	//FIM DOS OPERADORES LÓGICOS

	// UNÁRIOS (só interagem com uma variável por vez)

	numero := 10
	numero++     //vai acrescentar 1 unidade à variável numero
	numero += 15 //numero = numero + 15
	fmt.Println(numero)

	numero--     //vai decrementar em 1 unidade à variável numero
	numero -= 20 // numero = numero - 20

	numero *= 3 // numero = numero * 3

	numero /= 10 // numero = numero / 10

	numero %= 3 // numero = numero mod 3

	fmt.Println(numero)

	// FIM DOS OPERADOS UNÁRIOS

	// TERNÁRIOS (ex:. se o numero for maior do que 5, variável vai receber o texto "Maior que 5").
	// se o numero form menor do que 5, a variável vai recever o texto "Menor que 5"

	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}
	fmt.Println(texto)
}
