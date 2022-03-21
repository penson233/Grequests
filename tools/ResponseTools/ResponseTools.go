package ResponseTools

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)
//返回响应文本
func Transformresp(resp *http.Response) string{
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}
//返回json
func Jsondecode(resp string)  map[string]interface{}{
	r:=make(map[string]interface{})
	json.Unmarshal([]byte(resp),&r)
	return r
}
