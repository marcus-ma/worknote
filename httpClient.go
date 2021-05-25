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
	Query map[string]string
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
	//判断是否有body的数据
	if len(options.Payload)>0 {
		payload,_:=json.Marshal(options.Payload)
		body = bytes.NewReader(payload)
	}else {
		body = nil
	}

	req, _ := http.NewRequest(method, hc.baseUri+uri,body)
	
	//header请求头
	for key,value :=  range options.Header{
		req.Header.Set(key, value)
	}
	//query参数
	q := req.URL.Query()
	for key,value :=  range options.Query{
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	//请求发送
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return nil,err
	}
	data, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return data, err
}


func testDemo(){
	Client,_ := NewHttpClient(HttpClientOption{
		BaseUri:"https://api.baidu.com",
		TimeOut: time.Second*60,
	})
	postMaps := make(map[string]interface{})
	postMaps["body"]=map[string]interface{}{
		"realTimeRequestType":map[string]interface{}{
			"producttype":0,
			"startDate":revcData["date"].(string),
			"endDate":revcData["date"].(string),
			"levelOfDetails":2,
			"number":2000,
			"order":true,
			"performanceData":[]string{"cost"},
			"reportType":700,
			"statRange":2,//统计账户粒度
			"unitOfTime":5,//5按日
			"bstype":0,
		},
	}
	
	body,_:=Client.Request(http.MethodGet,apiUrl,HttpRequestOption{
		Payload: postMaps,
		Header: map[string]string{
			"Content-Type":"application/json;charset=UTF-8",
		},
	 	Query: map[string]string{
			"id":"1123",
		},
	})
	data := make(map[string]interface{})
	_ = json.Unmarshal(body, &data)
	fmt.Println(data)
}
