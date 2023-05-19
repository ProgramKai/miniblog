package wechat_sdk

import (
	"context"
	"sync"
	"time"
)

// WechatApp 微信应用
type WechatApp struct {
	*WechatApiClient
	// 应用 AgentId
	agentId int64
	// 应用 Secret
	corpSecret string
	// 接收消息 EncodingAESKey
	encodingAESKey string
	// 接收消息 Token
	token string
	//accessToken 企业微信access_token
	accessToken *token
	// tokenManager token管理器
	accessTokenManager TokenManager
}

func NewApp(baseURL, corpId string, agentId int64, corpSecret, encodingAESKey, appToken string, timeout time.Duration) *WechatApp {
	wechatApiClient := New(baseURL, corpId, timeout)
	app := &WechatApp{
		WechatApiClient: wechatApiClient,
		agentId:         agentId,
		corpSecret:      corpSecret,
		encodingAESKey:  encodingAESKey,
		token:           appToken,
		accessToken:     &token{mutex: &sync.RWMutex{}},
	}
	app.accessTokenManager = &token{
		mutex:        &sync.RWMutex{},
		getTokenFunc: app.GetToken,
	}
	return app
}

func (self *WechatApp) WithSaveCache(saveTokenCacheFunc SaveTokenCacheFunc) {
	self.accessTokenManager.(*token).saveTokenCacheFunc = saveTokenCacheFunc
}

func (self *WechatApp) WithGetTokenFromCacheFunc(getTokenFromCache GetTokenFormCacheFunc) {
	self.accessTokenManager.(*token).getTokenFormCacheFunc = getTokenFromCache
}

func (app *WechatApp) GetToken() (TokenInfo, error) {
	accessToken, err := app.getAccessToken(context.Background())
	if err != nil {
		return TokenInfo{}, err
	}
	return TokenInfo{
		Token:     accessToken.AccessToken,
		ExpiresIn: time.Duration(accessToken.ExpiresIn) * time.Second,
	}, nil
}
