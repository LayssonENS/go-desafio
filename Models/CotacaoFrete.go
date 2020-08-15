package models

// CalculaFreteRequest estrutura para cálculo de frete na api da frete rapido
type CalculaFreteRequest struct {
	Remetente        Remetente    `json:"remetente,omitempty"`
	Destinatario     Destinatario `json:"destinatario,omitempty"`
	Volumes          []Volume     `json:"volumes,omitempty"`
	CodigoPlataforma string       `json:"codigo_plataforma,omitempty"`
	Token            string       `json:"token,omitempty"`
}

// Destinatario informações do destinatario
type Destinatario struct {
	TipoPessoa        int64           `json:"tipo_pessoa,omitempty"`
	CnpjCpf           string          `json:"cnpj_cpf,omitempty"`
	InscricaoEstadual string          `json:"inscricao_estadual,omitempty"`
	Endereco          EnderecoEmpresa `json:"endereco,omitempty"`
}

// EnderecoEmpresa informações do endereco do destinatario
type EnderecoEmpresa struct {
	CEP string `json:"cep,omitempty"`
}

// Remetente informações do remetente do frete
type Remetente struct {
	CNPJ string `json:"cnpj,omitempty"`
}

// Volume informações dos itens que faz parte do frete
type Volume struct {
	Tipo           int64   `json:"tipo,omitempty"`
	Sku            string  `json:"sku,omitempty"`
	Descricao      string  `json:"descricao,omitempty"`
	Quantidade     int64   `json:"quantidade,omitempty"`
	Altura         float64 `json:"altura,omitempty"`
	Largura        float64 `json:"largura,omitempty"`
	Comprimento    float64 `json:"comprimento,omitempty"`
	Peso           float64 `json:"peso,omitempty"`
	Valor          float64 `json:"valor,omitempty"`
	VolumesProduto int64   `json:"volumes_produto,omitempty"`
}

// CotacaoResponse struct de calculo de frete na api da frete rapido
type CotacaoResponse struct {
	TokenOferta     string           `json:"token_oferta"`
	Transportadoras []Transportadora `json:"transportadoras"`
}

// Transportadora item da cotação
type Transportadora struct {
	Oferta           int64   `json:"oferta"`
	Cnpj             string  `json:"cnpj"`
	Logotipo         string  `json:"logotipo"`
	Nome             string  `json:"nome"`
	Servico          string  `json:"servico"`
	DescricaoServico string  `json:"descricao_servico"`
	PrazoEntrega     int64   `json:"prazo_entrega"`
	Validade         string  `json:"validade"`
	CustoFrete       float64 `json:"custo_frete"`
	PrecoFrete       float64 `json:"preco_frete"`
}
