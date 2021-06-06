package httpclient

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// HTTPClient make HTTP requests.
type HTTPClient struct {
	method  string
	url     string
	headers map[string]string
	body    map[string]string
	isForm  bool
	timeout time.Duration
}

// NewHTTPClient return a new HTTPClient.
func NewHTTPClient(method string, url string, headers map[string]string, body map[string]string, isForm bool, timeout time.Duration) *HTTPClient {
	return &HTTPClient{
		method:  method,
		url:     url,
		headers: headers,
		body:    body,
		isForm:  isForm,
		timeout: timeout,
	}
}

// SetURL reset url of the HTTPClient.
func (r *HTTPClient) SetURL(url string) {
	r.url = url
}

// SetParams append additional url info for HTTPClient.
func (r *HTTPClient) SetParams(params map[string]string) error {
	v := url.Values{}
	for key, value := range params {
		v.Add(key, value)
	}
	urlTemp, err := url.Parse(r.url)
	if err != nil {
		return err
	}
	urlTemp.RawQuery = v.Encode()
	r.url = urlTemp.String()
	return nil
}

// Do do HTTP request, return code, byte and error.
func (r *HTTPClient) Do() (int, []byte, error) {
	client := &http.Client{Timeout: r.timeout}
	var body []byte
	var err error
	if r.body != nil && !r.isForm {
		body, err = json.Marshal(r.body)
		if err != nil {
			return -1, nil, err
		}
	} else if r.body != nil && r.isForm {
		v := url.Values{}
		for key, value := range r.body {
			v.Add(key, value)
			body = []byte(v.Encode())
		}
	}
	req, err := http.NewRequest(r.method, r.url, bytes.NewReader(body))
	if err != nil {
		return -1, nil, err
	}
	if r.headers != nil {
		for k, v := range r.headers {
			req.Header.Add(k, v)
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		return -1, nil, err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1, nil, err
	}
	return resp.StatusCode, content, nil
}

// DoReturnResp do HTTP requests, return httpResp and error.
func (r *HTTPClient) DoReturnResp() (*http.Response, error) {
	client := new(http.Client)
	var body []byte
	var err error
	if r.body != nil && !r.isForm {
		body, err = json.Marshal(r.body)
		if err != nil {
			return nil, err
		}
	} else if r.body != nil && r.isForm {
		v := url.Values{}
		for key, value := range r.body {
			v.Add(key, value)
			body = []byte(v.Encode())
		}
	}
	req, err := http.NewRequest(r.method, r.url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	if r.headers != nil {
		for k, v := range r.headers {
			req.Header.Add(k, v)
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
