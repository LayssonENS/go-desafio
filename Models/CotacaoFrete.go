package Models

/*Struct dos campos de retorno da api*/
type ConsultaFrete struct {
	DataSituacao          string       `json:"data_situacao"`
	Nome                  string       `json:"nome"`
}