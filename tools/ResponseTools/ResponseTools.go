package ResponseTools

import (
	"encoding/json"
	"github.com/axgle/mahonia"
	"io/ioutil"
	"log"
	"net/http"
)

//返回响应文本
func Transformresp(resp *http.Response) string {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}

//对响应页面进行转码
func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

//返回json
func Jsondecode(resp string) map[string]interface{} {
	r := make(map[string]interface{})
	json.Unmarshal([]byte(resp), &r)
	return r
}
