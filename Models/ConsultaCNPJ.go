package Models

/*Struct dos campos de retorno da api*/
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

/*Billing*/
type Billing struct {
	Free     bool `json:"free"`
	Database bool `json:"database"`
}

/*Extra*/
type Extra struct {
}

/*Sócios*/
type Qsa struct {
	Qual string `json:"qual"`
	Nome string `json:"nome"`
}

/*Retorno JSON personalizado*/
type ResponsEempresa struct {
	Cnpj               string               `json:"cnpj"`
	UltimaAtualizacao  string               `json:"ultima_atualizacao"`
	Abertura           string               `json:"abertura"`
	Nome               string               `json:"nome"`
	Fantasia           string               `json:"fantasia"`
	Status             string               `json:"status"`
	Tipo               string               `json:"tipo"`
	Situacao           string               `json:"situacao"`
	CapitalSocial      string               `json:"capital_social"`
	Endereco           Endereco         	`json:"endereco"`
	Contato            Contato           	`json:"contato"`
	AtividadePrincipal AtividadePrincipal  `json:"atividade_principal"`
}

/*Atividade Principal*/
type AtividadePrincipal struct {
	Text string `json:"text"`
	Code string `json:"code"`
}

/*Contato*/
type Contato struct {
	Telefone string `json:"telefone"`
	Email    string `json:"email"`
}

/*Endereço da empresa*/
type Endereco struct {
	Bairro      string `json:"bairro"`
	Logradouro  string `json:"logradouro"`
	Numero      string `json:"numero"`
	Cep         string `json:"cep"`
	Municipio   string `json:"municipio"`
	Uf          string `json:"uf"`
	Complemento string `json:"complemento"`
}

/*Estrutura para mensagem com erro*/
type RespostaComErro struct {
	Erro string `json:"erro"`
}

