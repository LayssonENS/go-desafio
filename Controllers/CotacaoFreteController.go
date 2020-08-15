package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/LayssonENS/go-desafio/models"
)

//BuscarFrete sdas
func BuscarFrete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var responseObject models.CalculaFreteRequest
	_ = json.NewDecoder(r.Body).Decode(&responseObject)

	responseObject.Remetente = models.Remetente{CNPJ: "17184406000174"}
	responseObject.Token = "c8359377969ded682c3dba5cb967c07b"
	responseObject.CodigoPlataforma = "588604ab3"

	jsonValue, _ := json.Marshal(responseObject)

	/* Realiza consulta de cnpj na API: https://www.receitaws.com.br/v1/cnpj/  */
	res, err := http.Post("https://freterapido.com/api/external/embarcador/v1/quote-simulator", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.RespostaComErro{Erro: "Erro ao tentar realizar a consulta"})
		return
	}

	/* Passando Json pela Struct e verificando se teve erro */
	cotacoes, err := ioutil.ReadAll(res.Body)
	var response models.CotacaoResponse
	json.Unmarshal(cotacoes, &response)
	defer res.Body.Close()

	var cotacaot []Transportadora

	for _, item := range response.Transportadoras {
		cotacaot = append(cotacaot, Transportadora{
			Nome:         item.Nome,
			Servico:      item.Servico,
			PrazoEntrega: item.PrazoEntrega,
			PrecoFrete:   item.PrecoFrete,
		})
	}

	respostaJSON := CotacaoResponse{
		Transportadoras: cotacaot}
	//teste, _ := json.Marshal(CotacaoResponse)

	/*Retorna dados personalizados*/
	json.NewEncoder(w).Encode(respostaJSON)

}

// CotacaoResponse estrutura para resposta na rota de cotação
type CotacaoResponse struct {
	Transportadoras []Transportadora `json:"transportadoras"`
}

//Transportadora s
type Transportadora struct {
	Nome         string  `json:"nome"`
	Servico      string  `json:"servico"`
	PrazoEntrega int64   `json:"prazo_entrega"`
	PrecoFrete   float64 `json:"preco_frete"`
}
