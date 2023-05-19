package wechat_sdk

import (
	"context"
	"sync"
	"time"
)

// TokenManager token管理器
type TokenManager interface {
	Get(ctx context.Context) (TokenInfo, error)
	Refresh(ctx context.Context) (TokenInfo, error)
}

type GetTokenFormCacheFunc func(ctx context.Context, key string) (*TokenInfo, error)
type SaveTokenCacheFunc func(ctx context.Context, key string, tokenInfo *TokenInfo) error

// TokenInfo token信息
type TokenInfo struct {
	Token     string
	ExpiresIn time.Duration
}

var _ TokenManager = (*token)(nil)

// token 企业微信token
type token struct {
	mutex *sync.RWMutex
	TokenInfo
	lastRefresh           time.Time
	cacheKey              string
	getTokenFunc          GetTokenFunc
	getTokenFormCacheFunc GetTokenFormCacheFunc
	saveTokenCacheFunc    SaveTokenCacheFunc
}

// GetTokenFunc 获取token的函数
type GetTokenFunc func() (TokenInfo, error)

// SetTokenFunc 设置获取token的函数
func (t *token) SetTokenFunc(f GetTokenFunc) {
	t.getTokenFunc = f
}

// Get 获取token
func (t *token) Get(ctx context.Context) (TokenInfo, error) {
	t.mutex.RLock()
	if t.getTokenFormCacheFunc != nil {
		tokenInfo, err := t.getTokenFormCacheFunc(ctx, t.cacheKey)
		if err == nil && tokenInfo != nil {
			t.TokenInfo = *tokenInfo
		}
	}
	if t.Token != "" || time.Now().Sub(t.lastRefresh) >= t.ExpiresIn {
		t.mutex.RUnlock()
		_ = t.refreshToken()
		t.mutex.RLock()
	}
	defer t.mutex.RUnlock()
	return t.TokenInfo, nil
}

// refreshToken 刷新token
func (t *token) refreshToken() error {
	info, err := t.getTokenFunc()
	if err != nil {
		return err
	}
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.Token = info.Token
	t.ExpiresIn = info.ExpiresIn * time.Second
	t.lastRefresh = time.Now()
	if t.saveTokenCacheFunc != nil {
		_ = t.saveTokenCacheFunc(context.Background(), t.cacheKey, &t.TokenInfo)
	}
	return nil
}

// Refresh 刷新token
func (t *token) Refresh(ctx context.Context) (TokenInfo, error) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	err := t.refreshToken()
	if err != nil {
		return TokenInfo{}, err
	}
	return t.TokenInfo, nil
}
