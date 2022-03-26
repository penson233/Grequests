package RequestsTools

import (
	"math/rand"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

type Client struct {
	Proxies       map[string]string //代理
	Timeout       int               //超时时限
	RedirectCount int               //重定向的次数

}

//创建请求客户端
func (this *Client) CreateClient() *http.Client {
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

	jar, _ := cookiejar.New(nil)

	client = &http.Client{
		Transport: transport,
		Timeout:   time.Second * time.Duration(this.Timeout),
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if len(via) >= this.RedirectCount {
				return http.ErrUseLastResponse
			}
			return nil
		},
		Jar: jar,
	}

	return client
}
