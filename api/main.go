package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

//FUNÇÃO FOI UTILIZADA PARA GERAR A VARIÁVEL DE AMBIENTE COM A CHAVE ALEATÓRIA CONVERTIDA DE BASE
// 64 PARA STRING - COMO É FEITO UMA ÚNICA VEZ, O PROFESSOR APAGOU.
// MAS MANTENHO AQUI A TÍTULO DOCUMENTAL!
//func init() {
//	chave := make([]byte, 64)
//	if _, erro := rand.Read(chave); erro != nil {
//		log.Fatal(erro)
//	}
//	stringBase64 := base64.StdEncoding.EncodeToString(chave)
//	fmt.Println(stringBase64)
//}

func main() {

	config.Carregar()
	r := router.Gerar()

	//fmt.Println(config.SecretKey)

	fmt.Printf("Escutando na porta %d", config.Porta)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))

}
