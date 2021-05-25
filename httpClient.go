package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type HttpClientOption struct {
	BaseUri string
	TimeOut time.Duration

}
type HttpRequestOption struct {
	Payload map[string]interface{}
	Header map[string]string
}


type httpClient struct {
	baseUri string
	timeOut time.Duration
}

func NewHttpClient(options HttpClientOption) (*httpClient,error) {
	var(
		timeOut time.Duration
	)

	if options.BaseUri == "" {
		return nil,errors.New("baseUri must be had value")
	}

	if options.TimeOut==0 {
		timeOut = 30*time.Second
	}else {
		timeOut = options.TimeOut
	}
	return &httpClient{
		baseUri: options.BaseUri,
		timeOut: timeOut,
	},nil
}

//发起请求
func (hc *httpClient) Request(method string,uri string,options HttpRequestOption)(data []byte,err error) {

	var (
		body io.Reader
		resp *http.Response
	)
	if len(options.Payload)>0 {
		payload,_:=json.Marshal(options.Payload)
		body = bytes.NewReader(payload)
	}else {
		body = nil
	}

	req, _ := http.NewRequest(method, hc.baseUri+uri,body)
	for key,value :=  range options.Header{
		req.Header.Set(key, value)
	}


	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return nil,err
	}
	data, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return data, err

}