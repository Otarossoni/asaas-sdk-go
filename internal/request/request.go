package request

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"
)

type Headers map[string]interface{}
type QueryParams map[string]interface{}
type PathParams []interface{}

// Params Parâmetros para o método Request
type Params struct {
	Method       string
	URL          string
	Body         interface{}
	Headers      Headers
	Timeout      int
	PathParams   PathParams
	QueryParams  QueryParams
	BasicAuth    *BasicAuth
	HandleErrors *bool
}

// BasicAuth Usuário e senha usados na autenticação por BasicAuth
type BasicAuth struct {
	Username string
	Password string
}

// Response Retorno do método Request
type Response struct {
	StatusCode int
	Headers    Headers
	Body       map[string]interface{}
	RawBody    []byte
}

// New Efetua uma requsição http para uma API, microservice ou outro.
func New(params Params) (*Response, error) {
	var body *bytes.Reader

	// Verificando caso a requisição possua body então encodamos ele em JSON
	if params.Body != nil {
		data, err := json.Marshal(params.Body)

		if err != nil {
			return &Response{}, err
		}

		body = bytes.NewReader(data)
	}

	// Verificando caso a requisição possua PathParams então adicionamos na URL separados por /
	if len(params.PathParams) > 0 {
		params.URL = strings.TrimSuffix(params.URL, "/")

		for _, v := range params.PathParams {
			params.URL += "/" + toString(v)
		}

	}

	// Verificando caso a requisição possua QueryParams então adicionamos na URL
	if len(params.QueryParams) > 0 {
		query := url.Values{}

		for k, v := range params.QueryParams {
			query.Add(k, toString(v))
		}

		params.URL += "?" + query.Encode()
	}

	var request *http.Request
	var err error

	// Instanciando a requisição para depois o cliente pode executa-la.
	// Caso a requisição possua body então passamos a variavel do body, caso contrario passamos nil.
	// Não podemos passar a variavel body quando ela é nil porque se não da erro interno do GO por causa dos nils tipados e nils não tipados.
	if body == nil {
		request, err = http.NewRequest(params.Method, params.URL, nil)
	} else {
		request, err = http.NewRequest(params.Method, params.URL, body)
	}

	if err != nil {
		return &Response{}, err
	}

	if params.BasicAuth != nil {
		request.SetBasicAuth(params.BasicAuth.Username, params.BasicAuth.Password)
	}

	// Setando o header que indica que o body esta sendo enviado em formato JSON.
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	// Verificando se a requsição possui headers então setamos todos na requisição.
	if len(params.Headers) > 0 {

		for header, value := range params.Headers {
			request.Header.Set(header, toString(value))
		}

	}

	// Instanciando o client que ira executar a requsição.
	client := &http.Client{}

	// Verificando se algum timeout foi passado por parametro, caso não tenha sido passado então setamos 40 por padrão.
	if params.Timeout == 0 {
		params.Timeout = 40
	}

	// Setando o timeout no client.
	client.Timeout = time.Duration(params.Timeout) * time.Second

	// Executando a requisição.
	res, err := client.Do(request)

	if err != nil {
		return &Response{}, err
	}

	defer res.Body.Close()

	// Lendo os headers da resposta que veio da API.
	headers := Headers{}
	for name, values := range res.Header {
		headers[name] = values[0]
	}

	// Lendo a resposta veio da API.
	rawBody, err := io.ReadAll(res.Body)
	if err != nil {
		return &Response{StatusCode: res.StatusCode, Headers: headers, RawBody: rawBody}, err
	}

	// Decodificando a resposta da API.
	var untypedResponseBody interface{}
	err = json.Unmarshal(rawBody, &untypedResponseBody)

	// Verificando se a API retornou um Objeto JSON ou um Array de Objetos JSON.
	// Caso a API tenha retornado um Objeto de JSON então atribuimos direto no retorno, caso tiver retornar um Array ou outro tipo de dado,
	// então coloca o retorno da API dentro do campo "data" o retorno.
	var responseBody map[string]interface{}
	if jsonObject, isJsonObject := untypedResponseBody.(map[string]interface{}); isJsonObject {
		responseBody = jsonObject
	} else {
		responseBody = map[string]interface{}{"data": untypedResponseBody}
	}

	// Verificando se os errors devem ser tratados.
	if params.HandleErrors == nil || *params.HandleErrors {

		// Caso o StatusCode retornado seja maior que 299 então significa que deu algum erro.
		if res.StatusCode > 299 {
			return &Response{StatusCode: res.StatusCode, Headers: headers, RawBody: rawBody, Body: responseBody}, getError(responseBody, rawBody)
		}

	}

	/*
	 * Verificando se deu algum erro ao fazer o Unmarshal do JSON.
	 *
	 * Nós fazemos essa verificação aqui em baixo por que caso a API tenha retornado um erro (StatusCode > 299) e a essa resposta desse erro não
	 * tenha sido seja em JSON então a função Unmarshal() ira retornar erro de Parse do JSON, porém não podemos retornar logo de cara o erro de Parse
	 * do JSON porque pode ser que a API tenha retornar o erro em texto HTML ou PlainText, e nesse caso se a função HandleErrors estiver habilitada ela
	 * ira gerar um erro com o HTML da resposta ou com o texto plano da resposta.
	 *
	 * Além disso existem situações onde a API retorna apenas o StatusCode 200 (OK) ou 201 (Created), então também ira dar erro de Parse no JSON, porém esse erro
	 * é um erro falso, porque na verdade não deu erro nenhum, apenas a API não retornou nada, então tratamos isso arqui também.
	 */
	if err != nil {

		// Se não tiver body e for um StatusCode de sucesso (StatusCode < 300) então retorna uma response sem body
		if res.ContentLength == 0 && res.StatusCode < 300 {
			return &Response{StatusCode: res.StatusCode, Headers: headers, Body: map[string]interface{}{}, RawBody: []byte{}}, nil
		} else {
			return &Response{StatusCode: res.StatusCode, Headers: headers, RawBody: rawBody, Body: responseBody}, err
		}

	}

	return &Response{StatusCode: res.StatusCode, RawBody: rawBody, Body: responseBody, Headers: headers}, nil
}

// getError retorna o erro tratado de uma request.
func getError(body map[string]interface{}, rawBody ...[]byte) error {

	// Tentando buscar a mensagem de erro dentro do body.
	if body != nil {
		if body["error"] != nil {
			return errors.New(toString(body["error"]))
		} else if body["errors"] != nil {
			return errors.New(stringSlice(body["errors"])[0])
		} else if body["authErrors"] != nil {
			return errors.New(stringSlice(body["authErrors"])[0])
		} else if body["message"] != nil {
			return errors.New(toString(body["message"]))
		}
	}

	// Caso não tenha sido encontrada nenhuma mensagem de erro dentro do body então
	// verificamos se o rawBody foi passado por parametro, e retornamos todo o body em si.
	if len(rawBody) > 0 && len(rawBody[0]) > 0 {
		return errors.New(string(rawBody[0]))
	}

	// Caso não ache nada no body então retorna um erro genérico.
	return errors.New("Ocorreu uma falha ao realizar operação") //lint:ignore ST1005 ignore
}

// stringSlice convert []interface to []string
func stringSlice(itface interface{}) []string {
	s, ok := itface.([]interface{})
	if !ok {
		return []string{toString(itface)}
	}
	str := make([]string, len(s))
	for i, value := range s {
		str[i] = toString(value)
	}
	return str
}

// toString é o método responsável por retornar o valor de uma interface (pode ser ponteiro ou não) em string.
func toString(v interface{}) string {
	if v == nil {
		return ""
	}

	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {

		if rv.IsNil() {
			return ""
		} else {
			return fmt.Sprintf("%v", rv.Elem())
		}

	}

	return fmt.Sprintf("%v", rv)
}
