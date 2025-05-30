package asaas

// CustomerRequest é a struct usada para a criação de um novo Cliente na API Asaas
type CustomerRequest struct {
	Name                 string  `json:"name"`                 // Obrigatório - Nome do Cliente
	CpfCnpj              string  `json:"cpfCnpj"`              // Obrigatório - CPF ou CNPJ do Cliente
	Email                *string `json:"email"`                // E-mail do Cliente
	Phone                *string `json:"phone"`                // Telefone do Cliente
	MobilePhone          *string `json:"mobilePhone"`          // Número de telefone celular do Cliente
	Address              *string `json:"address"`              // Logradouro do endereço do Cliente
	AddressNumber        *string `json:"addressNumber"`        // Número do endereço do Cliente
	Complement           *string `json:"complement"`           // Complemento do endereço do Cliente
	Province             *string `json:"province"`             // Bairro do endereço do Cliente
	PostalCode           *string `json:"postalCode"`           // CEP do endereço do Cliente
	ExternalReference    *string `json:"externalReference"`    // Identificador do sistema integrado ao Asaas
	NotificationDisabled *bool   `json:"notificationDisabled"` // Realizar envio de notificações de cobrança ao cliente
	AdditionalEmails     *string `json:"additionalEmails"`     // E-mais adicionais, separados por ","
	MunicipalInscription *string `json:"municipalInscription"` // Inscrição Municipal do cliente
	StateInscription     *string `json:"stateInscription"`     // Inscrição Estadual do cliente
	Observations         *string `json:"observations"`         // Observações adicionais
	GroupName            *string `json:"groupName"`            // Nome do grupo ao qual o cliente pertence
	Company              *string `json:"company"`              // Empresa
	ForeignCustomer      *bool   `json:"foreignCustomer"`      // Define se o cliente é estrangeiro
}

// CustomerResponse é a struct usada para receber os dados da criação de um novo Cliente na API Asaas
type CustomerResponse struct {
	Object                string  `json:"object"`                // Tipo de recurso sendo criado
	Id                    string  `json:"id"`                    // ID do cliente na API Asaas
	Name                  string  `json:"name"`                  // Nome do Cliente
	Email                 *string `json:"email"`                 // E-mail do Cliente
	Company               *string `json:"company"`               // Empresa
	Phone                 *string `json:"phone"`                 // Telefone do Cliente
	MobilePhone           *string `json:"mobilePhone"`           // Número de telefone celular do Cliente
	Address               *string `json:"address"`               // Logradouro do endereço do Cliente
	AddressNumber         *string `json:"addressNumber"`         // Número do endereço do Cliente
	Complement            *string `json:"complement"`            // Complemento do endereço do Cliente
	Province              *string `json:"province"`              // Bairro do endereço do Cliente
	PostalCode            *string `json:"postalCode"`            // CEP do endereço do Cliente
	CpfCnpj               string  `json:"cpfCnpj"`               // CPF ou CNPJ do Cliente
	PersonType            string  `json:"personType"`            // Texto que descreve se o Cliente é pessoa Física ou Jurídica
	Deleted               bool    `json:"deleted"`               // Se o Cliente foi excluído na base de dados da API Asaas
	AdditionalEmails      *string `json:"additionalEmails"`      // E-mais adicionais, separados por ","
	ExternalReference     *string `json:"externalReference"`     // Identificador do sistema integrado ao Asaas
	Observations          *string `json:"observations"`          // Observações adicionais
	MunicipalInscription  *string `json:"municipalInscription"`  // Inscrição Municipal do cliente
	StateInscription      *string `json:"stateInscription"`      // Inscrição Estadual do cliente
	CanDelete             bool    `json:"canDelete"`             // Se o Cliente pode ser excluído
	CannotBeDeletedReason *string `json:"cannotBeDeletedReason"` // Motivo pelo qual o cliente não pode ser excluído
	CanEdit               bool    `json:"canEdit"`               // Se o Cliente pode ser editado
	CannotEditReason      *string `json:"cannotEditReason"`      // Motivo pelo qual o cliente não pode ser editado
	City                  *int    `json:"city"`                  // Código da cidade do Cliente
	CityName              *string `json:"cityName"`              // Nome da cidade do Cliente
	State                 *string `json:"state"`                 // UF do estado do Cliente
	Country               *string `json:"country"`               // Nome do país do Cliente
}

type ListCustomerResponse struct {
	Object     string             `json:"object"`     // Tipo de recurso sendo listado
	HasMore    bool               `json:"hasMore"`    // Flag que informa se há mais registros na lista
	TotalCount int                `json:"totalCount"` // Total de registros na lista
	Limit      int                `json:"limit"`      // Parâmetro "limit" da paginação
	Offset     int                `json:"offset"`     // Parâmetro "offset" da paginação
	Data       []CustomerResponse `json:"data"`       // Dados dos clientes encontrados para os filtros
}

// BillingRequest é a struct usada para a criação de uma nova Cobrança na API Asaas
type BillingRequest struct {
	Customer string `json:"customer"` // Obrigatório - ID do cliente gerado na API Asaas
	// Obrigatório - Forma de pagamento:
	//  * "UNDEFINED"
	//  * "BOLETO"
	//  * "CREDIT_CARD"
	//  * "PIX"
	BillingType                                string   `json:"billingType"`
	Value                                      float64  `json:"value"`                                      // Obrigatório - Valor da cobrança
	DueDate                                    string   `json:"dueDate"`                                    // Obrigatório - Data de vencimento da cobrança - Formato: yyyy-mm-dd
	Description                                *string  `json:"description"`                                // Descrição da cobrança (máx. 500 caracteres)
	DaysAfterDueDateToRegistrationCancellation *int     `json:"daysAfterDueDateToRegistrationCancellation"` // Dias após o vencimento para cancelamento do registro (somente para boleto bancário)
	ExternalReference                          *string  `json:"externalReference"`                          // Campo livre para busca
	InstallmentCount                           *int     `json:"installmentCount"`                           // Número de parcelas (somente no caso de cobrança parcelada)
	TotalValue                                 *float64 `json:"totalValue"`                                 // Informe o valor total de uma cobrança que será parcelada (somente no caso de cobrança parcelada). Caso enviado este campo o installmentValue não é necessário, o cálculo por parcela será automático
	InstallmentValue                           *float64 `json:"installmentValue"`                           // Valor de cada parcela (somente no caso de cobrança parcelada). Envie este campo em caso de querer definir o valor de cada parcela
	PostalService                              *bool    `json:"postalService"`                              // Define se a cobrança será enviada via Correios
}

// BillingResponse é a struct usada para receber os dados da criação de uma nova Cobrança na API Asaas
type BillingResponse struct {
	Object      string  `json:"object"`      // Tipo de recurso sendo criado
	Id          string  `json:"id"`          // ID da cobrança na API Asaas
	Customer    string  `json:"customer"`    // ID do cliente gerado na API Asaas
	DateCreated string  `json:"dateCreated"` // Data de criação da cobrança - Formato: yyyy-mm-dd
	Value       float64 `json:"value"`       // Valor da cobrança
	// Obrigatório - Forma de pagamento:
	//  * "UNDEFINED"
	//  * "BOLETO"
	//  * "CREDIT_CARD"
	//  * "PIX"
	BillingType           string     `json:"billingType"`
	CanBePaidAfterDueDate bool       `json:"canBePaidAfterDueDate"` // Informa se a cobrança pode ser paga após a data de vencimento
	CreditCard            CreditCard `json:"creditCard"`            // Informações do cartão de crédito usado no pagamento
	PixTransaction        string     `json:"pixTransaction"`        // ID da transação do PIX no caso de pagamento via PIX
	TransactionReceiptUrl string     `json:"transactionReceiptUrl"` // URL do recibo da transação
	PaymentDate           string     `json:"paymentDate"`           // Data em que o dinheiro irá cair na conta do asaas (no caso de cartão pode ser 30d) - Formato: yyyy-mm-dd
	ClientPaymentDate     string     `json:"clientPaymentDate"`     // Data em que o cliente efetuou o pagamento da cobrança - Formato: yyyy-mm-dd
	Status                string     `json:"status"`                // Situação da cobrança
	DueDate               string     `json:"dueDate"`               // Data de vencimento da cobrança - Formato: yyyy-mm-dd
	InvoiceUrl            string     `json:"invoiceUrl"`            // URL da cobrança, onde pode ser baixado o PDF / obtida a linha digitável / obtido o código do Boleto PIX
	InvoiceNumber         string     `json:"invoiceNumber"`         // Número da cobrança
	Deleted               bool       `json:"deleted"`               // Se a Cobrança foi excluída na base de dados da API Asaas
	NossoNumero           string     `json:"nossoNumero"`           // Campo "Nosso Número" do boleto
	BankSlipUrl           string     `json:"bankSlipUrl"`           // URL do boleto da cobrança
	PostalService         bool       `json:"postalService"`         // Informa se a cobrança foi enviada por e-mail
}

type CreditCard struct {
	CreditCardNumber string `json:"creditCardNumber"` // 4 digitos do número do cartão
	CreditCardBrand  string `json:"creditCardBrand"`  // Bandeira do cartão
}

type DeleteBillingResponse struct {
	Deleted bool   `json:"deleted"` // Se a Cobrança foi excluída
	Id      string `json:"id"`      // ID da Cobrança excluída
}

// SubscriptionRequest é a struct usada para a criação de uma nova assinatura na API Asaas
type SubscriptionRequest struct {
	CustomerId string `json:"customer"` // Obrigatório - Identificador único do cliente
	// Obrigatório - Forma de pagamento:
	//  * "UNDEFINED"
	//  * "BOLETO"
	//  * "CREDIT_CARD"
	//  * "PIX"
	BillingType string  `json:"billingType"`
	Value       float64 `json:"value"`       // Obrigatório - Valor da assinatura
	NextDueDate string  `json:"nextDueDate"` // Obrigatório - Vencimento da primeira cobrança
	// Obrigatório - Periodicidade da cobrança:
	//  * "WEEKLY"
	//  * "BIWEEKLY"
	//  * "MONTHLY"
	//  * "BIMONTHLY"
	//  * "QUARTERLY"
	//  * "SEMIANNUALLY"
	//  * "YEARLY"
	Cycle             string  `json:"cycle"`
	Description       *string `json:"description"`       // Descrição da assinatura (máx. 500 caracteres)
	EndDate           *string `json:"endDate"`           // Data limite para vencimento das cobranças
	MaxPayments       *int    `json:"maxPayments"`       // Número máximo de cobranças a serem geradas para esta assinatura
	ExternalReference *string `json:"externalReference"` // Identificador do sistema integrado ao Asaas
}

// SubscriptionResponse é a struct usada para receber os dados da criação de uma nova assinatura na API Asaas
type SubscriptionResponse struct {
	Object      string `json:"object"`      // Tipo de recurso sendo criado
	Id          string `json:"id"`          // ID da assinatura na API Asaas
	DateCreated string `json:"dateCreated"` // Data de criação da assinatura
	CustomerId  string `json:"customer"`    // Identificador único do cliente
	PaymentLink string `json:"paymentLink"` // Identificador único do link de pagamentos ao qual a assinatura pertence
	// Forma de pagamento:
	//  * "UNDEFINED"
	//  * "BOLETO"
	//  * "CREDIT_CARD"
	//  * "DEBIT_CARD"
	//  * "TRANSFER"
	//  * "DEPOSIT"
	//  * "PIX"
	BillingType string `json:"billingType"`
	// Obrigatório - Periodicidade da cobrança:
	//  * "WEEKLY"
	//  * "BIWEEKLY"
	//  * "MONTHLY"
	//  * "BIMONTHLY"
	//  * "QUARTERLY"
	//  * "SEMIANNUALLY"
	//  * "YEARLY"
	Cycle       string  `json:"cycle"`
	Value       float64 `json:"value"`       // Valor da assinatura
	NextDueDate string  `json:"nextDueDate"` // Vencimento do próximo pagamento a ser gerado
	EndDate     string  `json:"endDate"`     // Data limite para vencimento das cobranças
	Description string  `json:"description"` // Descrição da assinatura (máx. 500 caracteres)
	// Status da assinatura
	//  * "ACTIVE"
	//  * "EXPIRED"
	//  * "INACTIVE"
	Status            string `json:"status"`
	Deleted           bool   `json:"deleted"`           // Se a Assinatura foi excluído na base de dados da API Asaas
	MaxPayments       int    `json:"maxPayments"`       // Número máximo de cobranças a serem geradas para esta assinatura
	ExternalReference string `json:"externalReference"` // Identificador do sistema integrado ao Asaas
}

type ListSubscriptionResponse struct {
	Object     string                 `json:"object"`     // Tipo de recurso sendo listado
	HasMore    bool                   `json:"hasMore"`    // Flag que informa se há mais registros na lista
	TotalCount int                    `json:"totalCount"` // Total de registros na lista
	Limit      int                    `json:"limit"`      // Parâmetro "limit" da paginação
	Offset     int                    `json:"offset"`     // Parâmetro "offset" da paginação
	Data       []SubscriptionResponse `json:"data"`       // Dados das assinaturas encontradas para os filtros
}

type ListBillingResponse struct {
	Object     string            `json:"object"`     // Tipo de recurso sendo listado
	HasMore    bool              `json:"hasMore"`    // Flag que informa se há mais registros na lista
	TotalCount int               `json:"totalCount"` // Total de registros na lista
	Limit      int               `json:"limit"`      // Parâmetro "limit" da paginação
	Offset     int               `json:"offset"`     // Parâmetro "offset" da paginação
	Data       []BillingResponse `json:"data"`       // Dados dos pagamentos encontrados
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// ErrorResponse é a struct que é usada para receber os retornos de erro do Asaas
type ErrorResponse struct {
	Error   string `json:"error"`   // Slug do erro que retornou
	Message string `json:"message"` // Mensagem de erro relacionada ao campo
	Status  int    `json:"status"`  // Status/Código do erro
}
