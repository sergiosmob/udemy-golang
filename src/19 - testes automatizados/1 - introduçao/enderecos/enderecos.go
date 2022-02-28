package enderecos

import "strings"

//TipoDeEndereco verifica se um endereco tem um tipo válido e o retorna
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	enderecoEmLetraMinuscula := strings.ToLower(endereco)

	primeiraPalavraDoEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0] //" " é para dividir os nomes por ex. "RUA" "DOS" "BOBOS"

	enderecoTeUmTipoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTeUmTipoValido = true
		}
	}

	if enderecoTeUmTipoValido {
		return strings.Title(primeiraPalavraDoEndereco)
	}

	return "Tipo inválido"
}
