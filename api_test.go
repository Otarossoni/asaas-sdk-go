package asaas

import (
	"fmt"
	"testing"
)

// Testando criação de um Cliente
func TestSuccessOnCreateCustomer(t *testing.T) {

	response, asaasErr, err := CreateCustomer(CreateCustomerRequest{
		Name:    "PKG_TESTE3",
		CpfCnpj: "25373471000176",
	}, "$aact_MzkwODA2MWY2OGM3MWRlMDU2NWM3MzJlNzZmNGZhZGY6OjhmMmQzNGMwLWQ5NWMtNGEwYS05MDQ0LWNkMTdmZmE5ZTQ0YTo6JGFhY2hfNmZmZmY0NGYtYWE0Yi00Njk2LWEwZWYtZjA2ZDY0YzRkNzMw")

	if err != nil { // Erro interno
		t.Error("Erro inesperado!")
		t.Error(err.Error())

	} else if asaasErr != nil { // Erro da API Asaas
		t.Error("Erro não tratado Asaas!")
		t.Error(asaasErr.Message)
		t.Error(asaasErr.Status)
		t.Error(asaasErr.Error)

	} else { // Sucesso
		fmt.Print(response)
		t.Log(response.Id)
	}
}

// Testando a busca de um Cliente pelo ID Asaas
func TestSuccessOnGetCustomerById(t *testing.T) {

	response, asaasErr, err := GetCustomerByAsaasId(
		"cus_000006430156",
		"$aact_MzkwODA2MWY2OGM3MWRlMDU2NWM3MzJlNzZmNGZhZGY6OjhmMmQzNGMwLWQ5NWMtNGEwYS05MDQ0LWNkMTdmZmE5ZTQ0YTo6JGFhY2hfNmZmZmY0NGYtYWE0Yi00Njk2LWEwZWYtZjA2ZDY0YzRkNzMw",
	)

	if err != nil { // Erro interno
		t.Error("Erro inesperado!")
		t.Error(err.Error())

	} else if asaasErr != nil { // Erro da API Asaas
		t.Error("Erro não tratado Asaas!")
		t.Error(asaasErr.Message)
		t.Error(asaasErr.Status)
		t.Error(asaasErr.Error)

	} else { // Sucesso
		fmt.Print(response)
		t.Log(response.Id)
	}
}
