package core

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/go-resty/resty"
)

const (
	APIResponse        = "response"
	APIResponseHeaders = "headers"
	APIResponseCookies = "cookies"
)

type RestAPIWriterRequest struct {
	verb      string
	url       string
	urlParams map[string]interface{}
	headers   map[string]interface{}
	cookies   map[string]interface{}
	body      interface{}
	timeout   int
}

type RestAPIWriterResponse struct {
	StatusCode int
	Response   map[string]interface{}
}

type RestAPIWriter struct {
	transport *http.Transport
	logger    *logrus.Entry
}

func MakeAPIRequest() {
	transport := http.Transport{
		Dial: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 15 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
		MaxIdleConnsPerHost: 10,
	}
	client := resty.New().SetTransport(&transport).
		SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).SetTimeout(time.Duration(100*time.Millisecond) * time.Second)
	request := client.R()

	request.SetBody()
}

func (request *RestAPIWriterRequest) validate() error {
	if request.timeout == 0 {
		request.timeout = 30
	}

	return nil
}

func (request *RestAPIWriterRequest) WithURL(url string) *RestAPIWriterRequest {
	request.url = url
	return request
}

func (request *RestAPIWriterRequest) WithUrlParams(params map[string]interface{}) *RestAPIWriterRequest {
	request.urlParams = params
	return request
}

func (request *RestAPIWriterRequest) WithHeaders(headers map[string]interface{}) *RestAPIWriterRequest {
	request.headers = headers
	return request
}

func (request *RestAPIWriterRequest) WithBody(body interface{}) *RestAPIWriterRequest {
	request.body = body
	return request
}

func (request *RestAPIWriterRequest) WithVerb(verb string) *RestAPIWriterRequest {
	request.verb = verb
	return request
}

func (request *RestAPIWriterRequest) WithTimeout(timeout int) *RestAPIWriterRequest {
	request.timeout = timeout
	return request
}

func (response *RestAPIWriterResponse) Body() []byte {
	value := response.Response[APIResponse]
	if value != nil {
		return value.([]byte)
	}

	return []byte{}
}

func (writer *RestAPIWriter) processResponse(r *resty.Response) (*RestAPIWriterResponse, error) {
	writer.logger.Info("Processing response")
	response := &RestAPIWriterResponse{Response: map[string]interface{}{}}
	if r == nil {
		return response, nil
	}

	response.StatusCode = r.StatusCode()
	response.Response[APIResponse] = r.Body()
	response.Response[APIResponseHeaders] = r.Header()
	response.Response[APIResponseCookies] = r.Cookies()

	return response, nil
}
