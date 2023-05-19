package wechat_sdk

import (
	"log"
	"net/http"
	"time"
)

const (
	defaultBaseUrl = "https://qyapi.weixin.qq.com"
	defaultTimeout = time.Second * 15
)

// WechatApiClient 微信API客户端
type WechatApiClient struct {
	corpId     string
	httpClient *WechatApiHttpClient
}

func New(baseUrl, corpId string, timeout time.Duration) *WechatApiClient {
	if baseUrl == "" {
		baseUrl = defaultBaseUrl
	}
	if timeout == 0 {
		timeout = defaultTimeout
	}
	apiClient := &WechatApiClient{
		corpId: corpId,
		httpClient: &WechatApiHttpClient{
			baseUrl: baseUrl,
			Client: &http.Client{
				Timeout: timeout,
			},
		},
	}
	apiClient.httpClient.HttpBeforeRequestHook = apiClient.defaultHttpBeforeRequestHook()
	apiClient.httpClient.HttpAfterRequestHook = apiClient.defaultHttpAfterRequestHook()
	apiClient.httpClient.HttpOnErrorHook = apiClient.defaultHttpOnErrorHook()
	return apiClient
}

func (self *WechatApiClient) defaultHttpBeforeRequestHook() HttpBeforeRequestHook {
	return func(req *http.Request) {
		log.Println("defaultHttpBeforeRequestHook process")
	}
}

func (self *WechatApiClient) defaultHttpAfterRequestHook() HttpAfterRequestHook {
	return func(res *http.Response) {
		log.Println("defaultHttpAfterRequestHook process")
	}
}

func (self *WechatApiClient) defaultHttpOnErrorHook() HttpOnErrorHook {
	return func(req *http.Request, err error) {
		log.Println("defaultHttpOnErrorHook process")
	}
}
