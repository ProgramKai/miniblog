package wechat_sdk

import (
	"io"
	"net/http"
	"net/url"
)

type HttpBeforeRequestHook func(req *http.Request)
type HttpAfterRequestHook func(res *http.Response)
type HttpOnErrorHook func(req *http.Request, err error)

type WechatApiHttpClient struct {
	*http.Client
	baseUrl string
	HttpBeforeRequestHook
	HttpAfterRequestHook
	HttpOnErrorHook
}

func (self *WechatApiHttpClient) Do(req *http.Request) (*http.Response, error) {
	url, _ := url.Parse(self.baseUrl + req.URL.String())
	req.URL = url
	if self.HttpBeforeRequestHook != nil {
		self.HttpBeforeRequestHook(req)
	}
	res, err := self.Client.Do(req)
	if err != nil {
		if self.HttpOnErrorHook != nil {
			self.HttpOnErrorHook(req, err)
		}
		return nil, err
	}
	if self.HttpAfterRequestHook != nil {
		self.HttpAfterRequestHook(res)
	}
	return res, nil
}

func (self *WechatApiHttpClient) Get(url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, self.baseUrl, nil)
	if err != nil {
		return nil, err
	}
	return self.Do(req)
}

func (self *WechatApiHttpClient) Post(url string, contentType string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, self.baseUrl, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return self.Do(req)
}

func (self *WechatApiHttpClient) PostForm(url string, data map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, self.baseUrl, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	for k, v := range data {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()
	return self.Do(req)
}
