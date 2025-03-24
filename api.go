package asaas

import (
	"encoding/json"
	"os"

	"github.com/Otarossoni/asaas-sdk-go/internal/request"
)

type AsaasApi struct {
	BaseURL string
	Token   string
}

func NewAsaasApi(environment, token string) *AsaasApi {
	return &AsaasApi{
		BaseURL: getBaseURL(environment),
		Token:   getAccessToken(token),
	}
}

// CreateCustomer é o método responsável por realizar a criação do Cliente
func (a *AsaasApi) CreateCustomer(customerRequest CustomerRequest) (*CustomerResponse, *ErrorResponse, error) {

	params := request.Params{
		Method:  "POST",
		Body:    customerRequest,
		Headers: map[string]interface{}{"access_token": a.Token},
		URL:     a.BaseURL + "/v3/customers",
	}

	response, err := request.New(params)
	if err != nil {
		return nil, nil, err
	}

	if response.StatusCode > 300 {
		resp, err := parseError(response.RawBody)
		return nil, resp, err
	}

	var customerResponse CustomerResponse
	err = json.Unmarshal(response.RawBody, &customerResponse)
	return &customerResponse, nil, err
}

// GetCustomerByAsaasId é o método responsável por buscar um cliente pelo ID Asaas
func (a *AsaasApi) GetCustomerByAsaasId(customerId string) (*CustomerResponse, *ErrorResponse, error) {

	params := request.Params{
		Method:  "GET",
		Headers: map[string]interface{}{"access_token": a.Token},
		URL:     a.BaseURL + "/v3/customers/" + customerId,
	}

	response, err := request.New(params)
	if err != nil {
		return nil, nil, err
	}

	if response.StatusCode > 300 {
		resp, err := parseError(response.RawBody)
		return nil, resp, err
	}

	var customerResponse CustomerResponse
	err = json.Unmarshal(response.RawBody, &customerResponse)
	return &customerResponse, nil, err
}

// GetCustomerByCpfCnpj é o método responsável por buscar um cliente pelo CPF/CNPJ
func (a *AsaasApi) GetCustomerByCpfCnpj(customerCpfCnpj string) (*CustomerResponse, *ErrorResponse, error) {

	params := request.Params{
		Method:  "GET",
		Headers: map[string]interface{}{"access_token": a.Token},
		URL:     a.BaseURL + "/v3/customers",
		QueryParams: map[string]interface{}{
			"cpfCnpj": customerCpfCnpj,
		},
	}

	response, err := request.New(params)
	if err != nil {
		return nil, nil, err
	}

	if response.StatusCode > 300 {
		resp, err := parseError(response.RawBody)
		return nil, resp, err
	}

	var customerResponse ListCustomerResponse
	err = json.Unmarshal(response.RawBody, &customerResponse)

	if len(customerResponse.Data) > 0 {
		return &customerResponse.Data[0], nil, err
	}

	return &CustomerResponse{}, nil, err
}

// GetCustomerByName é o método responsável por buscar um cliente pelo nome
func (a *AsaasApi) GetCustomerByName(customerName string) (*CustomerResponse, *ErrorResponse, error) {

	params := request.Params{
		Method:  "GET",
		Headers: map[string]interface{}{"access_token": a.Token},
		URL:     a.BaseURL + "/v3/customers",
		QueryParams: map[string]interface{}{
			"name": customerName,
		},
	}

	response, err := request.New(params)
	if err != nil {
		return nil, nil, err
	}

	if response.StatusCode > 300 {
		resp, err := parseError(response.RawBody)
		return nil, resp, err
	}

	var customerResponse ListCustomerResponse
	err = json.Unmarshal(response.RawBody, &customerResponse)

	if len(customerResponse.Data) > 0 {
		return &customerResponse.Data[0], nil, err
	}

	return &CustomerResponse{}, nil, err
}

// CreateBilling é o método responsável por realizar a criação de uma Cobrança
func (a *AsaasApi) CreateBilling(billingRequest BillingRequest) (*BillingResponse, *ErrorResponse, error) {

	params := request.Params{
		Method:  "POST",
		Body:    billingRequest,
		Headers: map[string]interface{}{"access_token": a.Token},
		URL:     a.BaseURL + "/v3/payments",
	}

	response, err := request.New(params)
	if err != nil {
		return nil, nil, err
	}

	if response.StatusCode > 300 {
		resp, err := parseError(response.RawBody)
		return nil, resp, err
	}

	var billingResponse BillingResponse
	err = json.Unmarshal(response.RawBody, &billingResponse)
	return &billingResponse, nil, err
}

// GetBillingByAsaasId é o método responsável por buscar uma cobrança pelo ID Asaas
func (a *AsaasApi) GetBillingByAsaasId(billingId string) (*BillingResponse, *ErrorResponse, error) {

	params := request.Params{
		Method:  "GET",
		Headers: map[string]interface{}{"access_token": a.Token},
		URL:     a.BaseURL + "/v3/payments/" + billingId,
	}

	response, err := request.New(params)
	if err != nil {
		return nil, nil, err
	}

	if response.StatusCode > 300 {
		resp, err := parseError(response.RawBody)
		return nil, resp, err
	}

	var billingResponse BillingResponse
	err = json.Unmarshal(response.RawBody, &billingResponse)
	return &billingResponse, nil, err
}

// DeleteBilling é o método responsável por deletar uma cobrança pelo ID Asaas
func (a *AsaasApi) DeleteBilling(billingId string) (*DeleteBillingResponse, *ErrorResponse, error) {

	params := request.Params{
		Method:  "DELETE",
		Headers: map[string]interface{}{"access_token": a.Token},
		URL:     a.BaseURL + "/v3/payments/" + billingId,
	}

	response, err := request.New(params)
	if err != nil {
		return nil, nil, err
	}

	if response.StatusCode > 300 {
		resp, err := parseError(response.RawBody)
		return nil, resp, err
	}

	var deleteBillingResponse DeleteBillingResponse
	err = json.Unmarshal(response.RawBody, &deleteBillingResponse)
	return &deleteBillingResponse, nil, err
}

// CreateSubscription é o método responsável por realizar a criação de uma assinatura para um cliente
func (a *AsaasApi) CreateSubscription(subscriptionRequest SubscriptionRequest) (*SubscriptionResponse, *ErrorResponse, error) {

	params := request.Params{
		Method:  "POST",
		Body:    subscriptionRequest,
		Headers: map[string]interface{}{"access_token": a.Token},
		URL:     a.BaseURL + "/v3/subscriptions",
	}

	response, err := request.New(params)
	if err != nil {
		return nil, nil, err
	}

	if response.StatusCode > 300 {
		resp, err := parseError(response.RawBody)
		return nil, resp, err
	}

	var subscriptionResponse SubscriptionResponse
	err = json.Unmarshal(response.RawBody, &subscriptionResponse)
	return &subscriptionResponse, nil, err
}

// GetSubscriptionsByCustomerId é o método responsável por buscar as assinaturas de um cliente
func (a *AsaasApi) GetSubscriptionsByCustomerId(customerId string) ([]SubscriptionResponse, *ErrorResponse, error) {

	params := request.Params{
		Method:  "GET",
		Headers: map[string]interface{}{"access_token": a.Token},
		URL:     a.BaseURL + "/v3/subscriptions",
		QueryParams: map[string]interface{}{
			"customer": customerId,
		},
	}

	response, err := request.New(params)
	if err != nil {
		return nil, nil, err
	}

	if response.StatusCode > 300 {
		resp, err := parseError(response.RawBody)
		return nil, resp, err
	}

	var subscriptionResponse ListSubscriptionResponse
	err = json.Unmarshal(response.RawBody, &subscriptionResponse)
	return subscriptionResponse.Data, nil, err
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// getAccessToken é a função responsável por retornar o AccessToken do Asaas.
// Caso tenha sido passado um token por parâmetro pegamos o token passado por parâmetro, se não pegamos da variável de ambiente ASAAS_ACCESS_TOKEN.
func getAccessToken(asaasAccessToken ...string) string {
	if len(asaasAccessToken) > 0 {
		return asaasAccessToken[0]
	} else {
		return os.Getenv("ASAAS_ACCESS_TOKEN")
	}
}

// getBaseURL é a função responsável por validar o ambiente e a URL base.
func getBaseURL(environment string) string {
	if environment == "production" {
		return "https://api.asaas.com"
	}
	return "https://sandbox.asaas.com/api"
}

// parseError é a função que pega os dados do erro do Asaas e retorna em formato de Struct.
func parseError(body []byte) (*ErrorResponse, error) {
	var errResponse ErrorResponse
	if err := json.Unmarshal(body, &errResponse); err != nil {
		return nil, err
	}
	return &errResponse, nil
}
