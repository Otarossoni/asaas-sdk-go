package asaas

// CreateCustomerRequest é a struct usada para a criação de um novo Cliente na API Asaas
type CreateCustomerRequest struct {
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

// CreateCustomerResponse é a struct usada para receber os dados da criação de um novo Cliente na API Asaas
type CreateCustomerResponse struct {
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
	Object     string                   `json:"object"`     // Tipo de recurso sendo listado
	HasMore    bool                     `json:"hasMore"`    // Flag que informa se há mais registros na lista
	TotalCount int                      `json:"totalCount"` // Total de registros na lista
	Limit      int                      `json:"limit"`      // Parâmetro "limit" da paginação
	Offset     int                      `json:"offset"`     // Parâmetro "offset" da paginação
	Data       []CreateCustomerResponse `json:"data"`       // Dados dos clientes encontrados para os filtros
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// ErrorResponse é a struct que é usada para receber os retornos de erro do Asaas.
type ErrorResponse struct {
	Error   string `json:"error"`   // Slug do erro que retornou
	Message string `json:"message"` // Mensagem de erro relacionada ao campo
	Status  int    `json:"status"`  // Status/Código do erro
}
