package asaas

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

// Testando criação de um Cliente
func TestSuccessOnCreateCustomer(t *testing.T) {

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	asaasApi := NewAsaasApi(Sandbox, os.Getenv("ASAAS_ACCESS_TOKEN"))

	response, asaasErr, err := asaasApi.CreateCustomer(CustomerRequest{
		Name:    "TEST",
		CpfCnpj: "25373471000176",
	})

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

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	asaasApi := NewAsaasApi(Sandbox, os.Getenv("ASAAS_ACCESS_TOKEN"))

	response, asaasErr, err := asaasApi.GetCustomerByAsaasId("cus_000006587961")

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

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	asaasApi := NewAsaasApi(Sandbox, os.Getenv("ASAAS_ACCESS_TOKEN"))

	response, asaasErr, err := asaasApi.GetCustomerByCpfCnpj("25373471000176")

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

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	asaasApi := NewAsaasApi(Sandbox, os.Getenv("ASAAS_ACCESS_TOKEN"))

	response, asaasErr, err := asaasApi.GetCustomerByName("TESTE")

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

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	asaasApi := NewAsaasApi(Sandbox, os.Getenv("ASAAS_ACCESS_TOKEN"))

	response, asaasErr, err := asaasApi.CreateBilling(BillingRequest{
		Customer:    "cus_000006587961",
		BillingType: "BOLETO",
		Value:       101.99,
		DueDate:     "2025-03-25",
	})

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

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	asaasApi := NewAsaasApi(Sandbox, os.Getenv("ASAAS_ACCESS_TOKEN"))

	response, asaasErr, err := asaasApi.GetBillingByAsaasId("pay_8mmmgwqjsxlz7ttn")

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

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	asaasApi := NewAsaasApi(Sandbox, os.Getenv("ASAAS_ACCESS_TOKEN"))

	response, asaasErr, err := asaasApi.DeleteBilling("pay_8mmmgwqjsxlz7ttn")

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

// Testando criação de uma Assinatura
func TestSuccessOnCreateSubscription(t *testing.T) {

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	asaasApi := NewAsaasApi(Sandbox, os.Getenv("ASAAS_ACCESS_TOKEN"))

	response, asaasErr, err := asaasApi.CreateSubscription(SubscriptionRequest{
		CustomerId:  "cus_000006587961",
		BillingType: "PIX",
		Value:       5.01,
		NextDueDate: "2025-03-30",
		Cycle:       "MONTHLY",
	})

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

// Testando a busca de assinaturas de um cliente pelo seu ID Asaas
func TestSuccessOnGetSubscriptionsByCustomerId(t *testing.T) {

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	asaasApi := NewAsaasApi(Sandbox, os.Getenv("ASAAS_ACCESS_TOKEN"))

	response, asaasErr, err := asaasApi.GetSubscriptionsByCustomerId("cus_000006587961")

	if err != nil { // Erro interno
		t.Error("Erro inesperado!")
		t.Error(err.Error())

	} else if asaasErr != nil { // Erro da API Asaas
		t.Error("Erro não tratado Asaas!")
		t.Error(asaasErr.Message)
		t.Error(asaasErr.Status)
		t.Error(asaasErr.Error)

	} else { // Sucesso
		for _, item := range response {
			fmt.Print(item.Id)
			t.Log(item.Id)
		}
	}
}
