package models

/*DadosRequestCotacao - estrutura para dados recebidos via JSON POST*/
type DadosRequestCotacao struct {
	Remetente        Remetente    `json:"remetente"`
	Destinatario     Destinatario `json:"destinatario"`
	Volumes          []Volume     `json:"volumes"`
	CodigoPlataforma string       `json:"codigo_plataforma"`
	Token            string       `json:"token"`
}

/*Remetente - estrutura do remetente */
type Remetente struct {
	CNPJ string `json:"cnpj"`
}

/*Destinatario - estrutura do esdereco do destinatario */
type Destinatario struct {
	Endereco EnderecoEmpresa `json:"endereco"`
}

/*EnderecoEmpresa - Estrutura do endereco da empresa */
type EnderecoEmpresa struct {
	CEP string `json:"cep"`
}

/*Volume - estrutura dos dados dos volumes */
type Volume struct {
	Tipo           int64   `json:"tipo"`
	Sku            string  `json:"sku"`
	Descricao      string  `json:"descricao"`
	Quantidade     int64   `json:"quantidade"`
	Altura         float64 `json:"altura"`
	Largura        float64 `json:"largura"`
	Comprimento    float64 `json:"comprimento"`
	Peso           float64 `json:"peso"`
	Valor          float64 `json:"valor"`
	VolumesProduto int64   `json:"volumes_produto"`
}

/*CotacoesTransportadoras - estrutura das cotações  */
type CotacoesTransportadoras struct {
	Transportadoras []Transportadora `json:"transportadoras"`
}

/*Transportadora - estrutura dos destalhes das transportadoras*/
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

/*EstruturaCotacao  - Estrutura final para response*/
type EstruturaCotacao struct {
	Transportadoras []EstruturaTransportadora `json:"transportadoras"`
}

/*EstruturaTransportadora - detalhes personalizados das transportadoras */
type EstruturaTransportadora struct {
	Nome         string  `json:"nome"`
	Servico      string  `json:"servico"`
	PrazoEntrega int64   `json:"prazo_entrega"`
	PrecoFrete   float64 `json:"preco_frete"`
}
