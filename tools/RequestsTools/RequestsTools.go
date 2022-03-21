package RequestsTools

import (
	"Grequests/tools/ResponseTools"
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

type Requests struct {
	Url string
	Data string
	Headers map[string]string
	Proxies map[string]string
	File []string //{"file","1.php","<?php eval($_POST['penson'])?>","image/jpeg"}
	MutiData map[string]string

}


//上传文件
func (this *Requests)UploadFile() (*http.Request,error){
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


	fileWrite, err := bodyWrite.CreateFormFile(this.File[0], this.File[1],this.File[3])
	fileWrite.Write([]byte(this.File[2]))

	if err != nil {
		log.Println("err")
	}
	bodyWrite.Close() //要关闭，会将w.w.boundary刷写到w.writer中

	// 创建请求
	contentType := bodyWrite.FormDataContentType()
	req, err := http.NewRequest(http.MethodPost, "http://10.1.5.252", bodyBuf)
	req.Header.Set("Content-Type", contentType)
	return req,err

}

//创建请求客户端
func (this *Requests) createClient(method string) (*http.Client,*http.Request) {
	var client *http.Client

	if len(this.Proxies)==0 {
		//没有代理的情况
		client = &http.Client{}
	}else{
		//有代理的情况
		proxis :=make([]*url.URL,1)
		for _,value :=range this.Proxies{
			uri,_:=url.Parse(value)
			proxis=append(proxis, uri)
		}

		//随机选取代理ip
		rand.Seed(time.Now().Unix())
		proxy:=proxis[rand.Intn(len(proxis))]

		client =&http.Client{	Transport: &http.Transport{
			// 设置代理
			Proxy: http.ProxyURL(proxy),
		},}

	}


	var req *http.Request
	var err error


	if method=="GET" {
		req, err = http.NewRequest(method, this.Url,nil)
	}else if method=="POST"{
		if len(this.File) >0{
			req,err=this.UploadFile()
		}else{
			req,err =http.NewRequest("POST", this.Url, bytes.NewBuffer([]byte(this.Data)))
		}

	}
	for key,value :=range this.Headers{

		req.Header.Add(key,value)
	}



	if err != nil {
		fmt.Println(err)
	}


	return client,req
}

//GET 请求
func (this *Requests) Get() string  {

	resp:=this.request(http.MethodGet)
	return resp
	
}

//POST 请求
func (this *Requests) Post()string {

	resp:=this.request(http.MethodPost)
	return resp
}

func (this *Requests)request(method string) string{

	client,req:=this.createClient(method)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	return ResponseTools.Transformresp(resp)
}
