package asaas

import (
	"fmt"
	"testing"
)

// Testando criação de um Cliente
func TestSuccessOnCreateCustomer(t *testing.T) {

	response, asaasErr, err := CreateCustomer(CustomerRequest{
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
func TestSuccessOnGetCustomerByIdAsaas(t *testing.T) {

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

// Testando a busca de um Cliente pelo CPF/CNPJ
func TestSuccessOnGetCustomerByCpfCnpj(t *testing.T) {

	response, asaasErr, err := GetCustomerByCpfCnpj(
		"09567316000134",
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

// Testando a busca de um Cliente pelo Nome
func TestSuccessOnGetCustomerByName(t *testing.T) {

	response, asaasErr, err := GetCustomerByName(
		"TESTE",
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

// Testando criação de uma Cobrança
func TestSuccessOnCreateBilling(t *testing.T) {

	response, asaasErr, err := CreateBilling(BillingRequest{
		Customer:    "cus_000006430182",
		BillingType: "BOLETO",
		Value:       101.99,
		DueDate:     "2025-01-11",
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

// Testando a busca de uma Cobrança pelo ID Asaas
func TestSuccessOnGetBillingByIdAsaas(t *testing.T) {

	response, asaasErr, err := GetBillingByAsaasId(
		"pay_uir4qqv6v0844qch",
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

// Testando a exclusão de uma Cobrança pelo ID Asaas
func TestSuccessOnDeleteBilling(t *testing.T) {

	response, asaasErr, err := DeleteBilling(
		"pay_mko7zikl4k9xse1v",
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
