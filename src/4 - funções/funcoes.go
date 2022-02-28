package main

import "fmt"

func somar(n1 int8, n2 int8) int8 { // o tipo depois do parênteses é para definir o formato do retorno da função, caso tenha.
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) { //parâmetros do mesmo tipo -> podem ter o seu tipo decalrado apenas no úlimo.
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}
func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}
	resultado := f("Texto da função 1")
	fmt.Println(resultado)

	//resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	//fmt.Println(resultadoSoma, resultadoSubtracao)

	//resultadoSoma, _ := calculosMatematicos(10, 15)
	//fmt.Println(resultadoSoma)

	_, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSubtracao)

}
