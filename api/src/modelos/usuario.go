package modelos

import (
	"api/src/seguranca"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// Usuario representa um usuário utilizando a rede social Devbook
// as tabelas serão criadas na pasta "sql"
type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}

func (Usuario *Usuario) validar(etapa string) error {
	if Usuario.Nome == "" {
		return errors.New("O nome é obrigatório e não pode ser em branco")
	}

	if Usuario.Nick == "" {
		return errors.New("O nick é obrigatório e não pode ser em branco")
	}

	if Usuario.Email == "" {
		return errors.New("O Email é obrigatório e não pode ser em branco")
	}
	if erro := checkmail.ValidateFormat(Usuario.Email); erro != nil {
		return errors.New("O e-mail inserido é inválido")
	}

	if etapa == "cadastro" && Usuario.Senha == "" {
		return errors.New("A senha é obrigatória e não pode em branco")
	}

	return nil
}

//Preparar vai chamar os métodos para validar e formatar o usuário recebido
func (Usuario *Usuario) Preparar(etapa string) error {
	if erro := Usuario.validar(etapa); erro != nil {
		return erro
	}
	if erro := Usuario.formatar(etapa); erro != nil {
		return erro
	}
	return nil
}

func (Usuario *Usuario) formatar(etapa string) error {
	Usuario.Nome = strings.TrimSpace(Usuario.Nome)
	Usuario.Nick = strings.TrimSpace(Usuario.Nick)
	Usuario.Email = strings.TrimSpace(Usuario.Email)

	if etapa == "cadastro" {

		senhaComHash, erro := seguranca.Hash(Usuario.Senha)

		if erro != nil {
			return erro
		}

		Usuario.Senha = string(senhaComHash)

	}

	return nil
}
