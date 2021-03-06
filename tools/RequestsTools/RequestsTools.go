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
	Data          string
	Params        map[string]string
	File          []string //{"file","1.php","<?php eval($_POST['penson'])?>","image/jpeg"}
	MutiData      map[string]string
	Json          string
	Client        *http.Client
	Headers       map[string]string
	Proxies       map[string]string
	Timeout       int //超时时限
	RedirectCount int //重定向的次数
}

//上传文件
func (this *Requests) UploadFile(urlink string) (*http.Request, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWrite := NewWriter(bodyBuf)
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

//GET 请求
func (this *Requests) Get(url string) *http.Response {
	resp := this.request(http.MethodGet, url)
	return resp

}

//POST 请求
func (this *Requests) Post(url string) *http.Response {

	resp := this.request(http.MethodPost, url)
	return resp
}

func (this *Requests) handelreq(method string, link string, req *http.Request, err error) *http.Request {
	//处理get参数
	var urlink string
	params := ""
	if len(this.Params) != 0 {
		for key, value := range this.Params {
			params += key + "=" + url.QueryEscape(value) + "&"
		}
		params = params[:len(params)-1]
		urlink = link + "?" + params

	} else {
		urlink = link
	}

	if method == "GET" {

		req, err = http.NewRequest(method, urlink, nil)

	} else if method == "POST" {

		if len(this.File) > 0 {
			req, err = this.UploadFile(urlink)
		} else {

			//处理data
			if len(this.Json) != 0 {
				if len(this.Headers) <= 0 {
					this.Headers = make(map[string]string, 0)
				}
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
				if len(this.Headers) <= 0 {
					this.Headers = make(map[string]string, 0)
					this.Headers["Content-Type"] = "application/x-www-form-urlencoded"
				}

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
	return req
}

//创建客户端
func (this *Requests) createClient() *http.Client {

	var client *http.Client

	var transport = &http.Transport{}

	if len(this.Proxies) > 0 {
		//有代理的情况
		proxis := make([]*url.URL, 1)
		for _, value := range this.Proxies {
			uri, _ := url.Parse(value)
			proxis = append(proxis, uri)
		}

		//随机选取代理ip
		rand.Seed(time.Now().Unix())
		proxy := proxis[rand.Intn(len(proxis))]

		transport = &http.Transport{
			// 设置代理
			Proxy: http.ProxyURL(proxy),
		}
	}

	if this.Timeout == 0 {
		this.Timeout = 10
	}
	if this.RedirectCount < 2 {
		this.RedirectCount = 2
	}

	//jar, _ := cookiejar.New(nil)

	client = &http.Client{
		Transport: transport,
		Timeout:   time.Second * time.Duration(this.Timeout),
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if len(via) >= this.RedirectCount {
				return http.ErrUseLastResponse
			}
			return nil
		},
		//Jar: jar,
	}

	return client
}

func (this *Requests) request(method string, link string) *http.Response {
	client := this.createClient()

	if this.Client != nil {
		client = this.Client
	}

	var req *http.Request
	var err error

	//处理请求
	req = this.handelreq(method, link, req, err)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	return resp
}

func (this *Requests) Text(resp *http.Response) string {
	return ResponseTools.Transformresp(resp)
}
