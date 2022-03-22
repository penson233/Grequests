package RequestsTools

import (
	"bytes"
	"fmt"
	"github.com/penson233/Grequests/tools/ResponseTools"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

type Requests struct {
	Url      string
	Data     string
	Params   map[string]string
	Headers  map[string]string
	Proxies  map[string]string
	File     []string //{"file","1.php","<?php eval($_POST['penson'])?>","image/jpeg"}
	MutiData map[string]string
	Json     string
}

//上传文件
func (this *Requests) UploadFile(urlink string) (*http.Request, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWrite := NewWriter(bodyBuf)
	//file, err := os.Open(this.File[2])
	//defer file.Close()
	//if err != nil {
	//	log.Println("err")
	//}
	for key, val := range this.MutiData {
		_ = bodyWrite.WriteField(key, val)
	}

	fileWrite, err := bodyWrite.CreateFormFile(this.File[0], this.File[1], this.File[3])
	fileWrite.Write([]byte(this.File[2]))

	if err != nil {
		log.Println("err")
	}
	bodyWrite.Close() //要关闭，会将w.w.boundary刷写到w.writer中

	// 创建请求
	contentType := bodyWrite.FormDataContentType()
	req, err := http.NewRequest(http.MethodPost, urlink, bodyBuf)
	req.Header.Set("Content-Type", contentType)
	return req, err

}

//创建请求客户端
func (this *Requests) createClient(method string) (*http.Client, *http.Request) {
	var client *http.Client

	if len(this.Proxies) == 0 {
		//没有代理的情况
		client = &http.Client{}
	} else {
		//有代理的情况
		proxis := make([]*url.URL, 1)
		for _, value := range this.Proxies {
			uri, _ := url.Parse(value)
			proxis = append(proxis, uri)
		}

		//随机选取代理ip
		rand.Seed(time.Now().Unix())
		proxy := proxis[rand.Intn(len(proxis))]

		client = &http.Client{Transport: &http.Transport{
			// 设置代理
			Proxy: http.ProxyURL(proxy),
		}}

	}

	var req *http.Request
	var err error

	//处理get参数
	var urlink string
	params := ""
	if len(this.Params) != 0 {
		for key, value := range this.Params {
			params += key + "=" + url.QueryEscape(value) + "&"
		}
		params = params[:len(params)-1]
		urlink = this.Url + "?" + params

	} else {
		urlink = this.Url
	}

	if method == "GET" {

		req, err = http.NewRequest(method, urlink, nil)

	} else if method == "POST" {

		if len(this.File) > 0 {
			req, err = this.UploadFile(urlink)
		} else {

			//处理data
			_, ok := this.Headers["Content-Type"]
			if len(this.Json) != 0 || ok {
				this.Headers["Content-Type"] = "application/json"
				req, err = http.NewRequest("POST", urlink, bytes.NewBuffer([]byte(this.Json)))
			} else {
				//param:=""
				//if len(this.Data)!=0{
				//	for key,value :=range this.Data{
				//		param+=key+"="+url.QueryEscape(value)+"&"
				//	}
				//}
				//param=param[:len(param)-1]

				req, err = http.NewRequest("POST", urlink, bytes.NewBuffer([]byte(this.Data)))
			}

		}

	}
	for key, value := range this.Headers {

		req.Header.Add(key, value)
	}

	if err != nil {
		fmt.Println(err)
	}

	return client, req
}

//GET 请求
func (this *Requests) Get() string {

	resp := this.request(http.MethodGet)
	return resp

}

//POST 请求
func (this *Requests) Post() string {

	resp := this.request(http.MethodPost)
	return resp
}

func (this *Requests) request(method string) string {

	client, req := this.createClient(method)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	return ResponseTools.Transformresp(resp)
}
