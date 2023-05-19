package wechat_sdk

import (
	"cn.xdmnb/study/miniblog/pkg/wechat_sdk/domain"
	"context"
	"github.com/goccy/go-json"
	"io/ioutil"
	"net/http"
	"net/url"
)

// WechatCompanyBasicApi 企业微信基础接口
type WechatCompanyBasicApi interface {
	// GetAccessToken 获取企业微信的access_token
	getAccessToken(ctx context.Context) (*domain.AccessTokenResp, error)
}

func (app *WechatApp) getAccessToken(ctx context.Context) (*domain.AccessTokenResp, error) {
	url, _ := url.Parse("/cgi-bin/gettoken")
	query := url.Query()
	query.Set("corpid", app.WechatApiClient.corpId)
	query.Set("corpsecret", app.corpSecret)
	encode := query.Encode()
	url.RawQuery = encode
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, err
	}
	res, err := app.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	resBodyBytes, err := ioutil.ReadAll(res.Body)
	result := &domain.AccessTokenResp{}
	json.Unmarshal(resBodyBytes, result)
	return result, nil
}
