package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/LayssonENS/go-desafio/models"
	"github.com/Nhanderu/brdoc"
	"github.com/gorilla/mux"
)

/*BuscarCnpj Função para retornar dados da empresa */
func BuscarCnpj(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	/*Pegando parâmetro passado via GET*/
	cnpj := mux.Vars(r)

	/* Valida se Cnpj é válido e retorna caso não for*/
	if !brdoc.IsCNPJ(cnpj["cnpj"]) {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(models.RespostaComErro{Erro: "Cnpj Invalido"})
		return
	}

	/* Realiza consulta de cnpj na API: https://www.receitaws.com.br/v1/cnpj/  */
	result, err := http.Get("https://www.receitaws.com.br/v1/cnpj/" + cnpj["cnpj"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.RespostaComErro{Erro: "Erro ao tentar realizar a consulta"})
		return
	}

	/* Passando Json pela Struct e verificando se teve erro */
	responseConsulta, err := ioutil.ReadAll(result.Body)
	var responseObject models.ConsultaCnpj
	json.Unmarshal(responseConsulta, &responseObject)

	/*JSON Personalizado*/
	respostaJSON := models.DadosEmpresa{Empresa: models.ResponseEmpresa{Cnpj: responseObject.Cnpj,
		UltimaAtualizacao: responseObject.UltimaAtualizacao,
		Abertura:          responseObject.Abertura,
		Nome:              responseObject.Nome,
		Fantasia:          responseObject.Fantasia,
		Status:            responseObject.Status,
		Tipo:              responseObject.Tipo,
		Situacao:          responseObject.Situacao,
		CapitalSocial:     responseObject.CapitalSocial,
		Endereco: models.Endereco{
			Bairro:      responseObject.Bairro,
			Logradouro:  responseObject.Logradouro,
			Numero:      responseObject.Numero,
			Cep:         responseObject.Cep,
			Municipio:   responseObject.Municipio,
			Uf:          responseObject.Uf,
			Complemento: responseObject.Complemento,
		},
		Contato: models.Contato{
			Telefone: responseObject.Telefone,
			Email:    responseObject.Email,
		},
		AtividadePrincipal: models.AtividadePrincipal{
			Text: responseObject.AtividadePrincipal[0].Text,
			Code: responseObject.AtividadePrincipal[0].Code,
		},
	},
	}

	/*Retorna dados personalizados*/
	json.NewEncoder(w).Encode(respostaJSON)

}
