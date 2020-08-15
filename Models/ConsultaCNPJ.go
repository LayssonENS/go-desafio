package models

/*ConsultaCnpj Struct dos campos de retorno da api*/
type ConsultaCnpj struct {
	DataSituacao          string       `json:"data_situacao"`
	Nome                  string       `json:"nome"`
	Uf                    string       `json:"uf"`
	Telefone              string       `json:"telefone"`
	Email                 string       `json:"email"`
	Situacao              string       `json:"situacao"`
	Bairro                string       `json:"bairro"`
	Logradouro            string       `json:"logradouro"`
	Numero                string       `json:"numero"`
	Cep                   string       `json:"cep"`
	Municipio             string       `json:"municipio"`
	Abertura              string       `json:"abertura"`
	NaturezaJuridica      string       `json:"natureza_juridica"`
	Cnpj                  string       `json:"cnpj"`
	UltimaAtualizacao     string       `json:"ultima_atualizacao"`
	Status                string       `json:"status"`
	Tipo                  string       `json:"tipo"`
	Fantasia              string       `json:"fantasia"`
	Complemento           string       `json:"complemento"`
	Efr                   string       `json:"efr"`
	MotivoSituacao        string       `json:"motivo_situacao"`
	SituacaoEspecial      string       `json:"situacao_especial"`
	DataSituacaoEspecial  string       `json:"data_situacao_especial"`
	CapitalSocial         string       `json:"capital_social"`
	Message               string       `json:"message"`
	AtividadePrincipal    []Atividades `json:"atividade_principal"`
	AtividadesSecundarias []Atividades `json:"atividades_secundarias"`
	Qsa                   []Qsa        `json:"qsa"`
	Extra                 Extra        `json:"extra"`
	Billing               Billing      `json:"billing"`
}

/*Atividades da empresa*/
type Atividades struct {
	Text string `json:"text"`
	Code string `json:"code"`
}

/*Billing estrutura*/
type Billing struct {
	Free     bool `json:"free"`
	Database bool `json:"database"`
}

/*Extra estrutura*/
type Extra struct {
}

/*Qsa estrutura*/
type Qsa struct {
	Qual string `json:"qual"`
	Nome string `json:"nome"`
}

/*ResponseEmpresa Retorno JSON personalizado*/
type ResponseEmpresa struct {
	Cnpj               string             `json:"cnpj"`
	UltimaAtualizacao  string             `json:"ultima_atualizacao"`
	Abertura           string             `json:"abertura"`
	Nome               string             `json:"nome"`
	Fantasia           string             `json:"fantasia"`
	Status             string             `json:"status"`
	Tipo               string             `json:"tipo"`
	Situacao           string             `json:"situacao"`
	CapitalSocial      string             `json:"capital_social"`
	Endereco           Endereco           `json:"endereco"`
	Contato            Contato            `json:"contato"`
	AtividadePrincipal AtividadePrincipal `json:"atividade_principal"`
}

/*AtividadePrincipal estrutura*/
type AtividadePrincipal struct {
	Text string `json:"text"`
	Code string `json:"code"`
}

/*Contato estrutura*/
type Contato struct {
	Telefone string `json:"telefone"`
	Email    string `json:"email"`
}

/*Endereco da empresa*/
type Endereco struct {
	Bairro      string `json:"bairro"`
	Logradouro  string `json:"logradouro"`
	Numero      string `json:"numero"`
	Cep         string `json:"cep"`
	Municipio   string `json:"municipio"`
	Uf          string `json:"uf"`
	Complemento string `json:"complemento"`
}

/*RespostaComErro estrutura de erro*/
type RespostaComErro struct {
	Erro string `json:"erro"`
}
