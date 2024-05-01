package tax

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type RequestBody struct {
	time           int
	signature      string
	signatureKeyId string
	packets        any // use for send invoice
	packet         any
}

type Request struct {
	Url      string
	Path     string
	Method   string
	Body     RequestBody
	Header   map[string]any
	Priority string
	Sync     bool
}

type RequestMethods interface {
	MakeHeader(string)
	MakeUrl(string)
	MakeBody(*Packet, string, string, int)
	Send(any) (*[]byte, int, error)
}

func (r *Request) MakeHeader(token string) {
	headers := map[string]any{
		"requestTraceId": uuid.New().String(),
	}
	_timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	if token != "" {
		headers["Authorization"] = fmt.Sprintf("Bearer %s", token)
	}
	headers["timestamp"] = _timestamp
	r.Header = headers
}

func (r *Request) MakeUrl(TaxUrl string) {
	if r.Sync {
		TaxUrl += "/sync"
	} else {
		if r.Priority == "" {
			r.Priority = "normal-enqueue"
		}
		TaxUrl += fmt.Sprintf("/async/%s", r.Priority)
	}
	r.Url = fmt.Sprintf("%s/%s", TaxUrl, r.Path)
}

func (r *Request) MakeBody(body interface{}, signature string, signatureKeyId string, t int) {
	result := RequestBody{
		signature:      signature,
		signatureKeyId: signatureKeyId,
	}
	if t == 0 {
		t = 1
	}
	result.time = t
	switch body.(type) {
	case map[string]interface{}:
		result.packet = body
	case []map[string]interface{}:
		result.packets = body
	}
	r.Body = result
}

func (r *Request) Send(responseType any) (any, int, error) {
	status := 400
	jsonBody, err := json.Marshal(r.Body)
	if err != nil {
		return nil, 400, err
	}
	bodyBytesReader := bytes.NewReader(jsonBody)
	request, err := http.NewRequest(strings.ToUpper(strings.ToUpper(r.Method)), r.Url, bodyBytesReader)
	if err != nil {
		return nil, status, err
	}
	request.Header.Set("Content-Type", "application/json")
	for key, val := range r.Header {
		request.Header.Set(key, val.(string))
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := http.Client{Transport: tr}
	response, err := client.Do(request)
	if err != nil {
		return nil, status, err
	}
	defer response.Body.Close()
	if err != nil {
		return nil, status, err
	}
	responseByte, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, status, err
	}
	status = response.StatusCode
	err = json.Unmarshal(responseByte, &responseType)
	//rtType := responseType.(reflect.TypeOf(responseType))
	return responseType, status, err
}
