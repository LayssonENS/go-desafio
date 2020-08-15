package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/LayssonENS/go-desafio/models"
)

/*BuscarFrete - função para cotação do frete */
func BuscarFrete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	/* Recebendo dados enviados via post para montar JSON  */
	var responseObject models.DadosRequestCotacao
	err := json.NewDecoder(r.Body).Decode(&responseObject)

	/* Preenchendo struct par consulta em api */
	responseObject.Remetente = models.Remetente{CNPJ: "17184406000174"}
	responseObject.Token = "c8359377969ded682c3dba5cb967c07b"
	responseObject.CodigoPlataforma = "588604ab3"

	urlCotacao := "https://freterapido.com/api/external/embarcador/v1/quote-simulator"

	/* Valida se os dados recebidos */
	erroValidacao := validarDados(responseObject)
	if len(erroValidacao) > 0 {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(erroValidacao)
		return
	}

	dadosRecebidosJSON, _ := json.Marshal(responseObject)

	/* Realizando Consulta  */
	result, err := http.Post(urlCotacao, "application/json", bytes.NewBuffer(dadosRecebidosJSON))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.RespostaComErro{Erro: "Erro ao tentar realizar a consulta"})
		return
	}

	/* Recebendo resultado e preenchendo struct */
	cotacoes, err := ioutil.ReadAll(result.Body)
	var response models.CotacoesTransportadoras
	json.Unmarshal(cotacoes, &response)
	defer result.Body.Close()

	var transportadorasCotacoes []models.EstruturaTransportadora

	/* Montando struct com dados recebidos via JSON  */
	for _, item := range response.Transportadoras {
		transportadorasCotacoes = append(transportadorasCotacoes, models.EstruturaTransportadora{
			Nome:         item.Nome,
			Servico:      item.Servico,
			PrazoEntrega: item.PrazoEntrega,
			PrecoFrete:   item.PrecoFrete,
		})
	}

	/* Montando struct para resposta das cotações  */
	respostaJSON := models.EstruturaCotacao{
		Transportadoras: transportadorasCotacoes}

	/* Retorna dados personalizados*/
	json.NewEncoder(w).Encode(respostaJSON)

}

/*Função usada para validar dados recebidos */
func validarDados(dados models.DadosRequestCotacao) string {

	if len(dados.Destinatario.Endereco.CEP) != 8 {
		return "CEP Inválido"
	}
	if dados.Volumes[0].Tipo < 0 {
		return "Tipo Inválido"
	}
	if dados.Volumes[0].Quantidade == 0 {
		return "Quantidade Inválida"
	}
	if dados.Volumes[0].Peso == 0 {
		return "Peso Inválido"
	}
	if dados.Volumes[0].Valor == 0 {
		return "Valor Inválido"
	}
	if dados.Volumes[0].Altura == 0 {
		return "Altura Inválida"
	}
	if dados.Volumes[0].Largura == 0 {
		return "Largura Inválida"
	}
	if dados.Volumes[0].Comprimento == 0 {
		return "Tipo Inválido"
	}

	return ""
}
