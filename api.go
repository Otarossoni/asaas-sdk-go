package asaas

import (
	"encoding/json"
	"os"

	"github.com/Otarossoni/asaas-sdk-go/internal/request"
)

// Sandbox = "https://sandbox.asaas.com/api"
// Produção = "https://api.asaas.com"
const BASE_URL = "https://sandbox.asaas.com/api"

// CreateCustomer é o método responsável por realizar a criação do Cliente
func CreateCustomer(createCustomerRequest CreateCustomerRequest, asaasAccesstoken ...string) (*CreateCustomerResponse, *ErrorResponse, error) {

	params := request.Params{
		Method:  "POST",
		Body:    createCustomerRequest,
		Headers: map[string]interface{}{"access_token": getAccessToken(asaasAccesstoken...)},
		URL:     BASE_URL + "/v3/customers",
	}

	response, err := request.New(params)
	if err != nil {
		return nil, nil, err
	}

	if response.StatusCode > 300 {
		resp, err := parseError(response.RawBody)
		return nil, resp, err
	}

	var customerResponse CreateCustomerResponse
	err = json.Unmarshal(response.RawBody, &customerResponse)
	return &customerResponse, nil, err
}

// GetCustomerByAsaasId é o método responsável por buscar um cliente pelo ID Asaas
func GetCustomerByAsaasId(customerId string, asaasAccesstoken ...string) (*CreateCustomerResponse, *ErrorResponse, error) {

	params := request.Params{
		Method:  "GET",
		Headers: map[string]interface{}{"access_token": getAccessToken(asaasAccesstoken...)},
		URL:     BASE_URL + "/v3/customers/" + customerId,
	}

	response, err := request.New(params)
	if err != nil {
		return nil, nil, err
	}

	if response.StatusCode > 300 {
		resp, err := parseError(response.RawBody)
		return nil, resp, err
	}

	var customerResponse CreateCustomerResponse
	err = json.Unmarshal(response.RawBody, &customerResponse)
	return &customerResponse, nil, err
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

// parseError é a função que pega os dados do erro do Asaas e retorna em formato de Struct.
func parseError(body []byte) (*ErrorResponse, error) {
	var errResponse ErrorResponse
	if err := json.Unmarshal(body, &errResponse); err != nil {
		return nil, err
	}
	return &errResponse, nil
}
