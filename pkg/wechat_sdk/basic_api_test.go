package wechat_sdk

import (
	"testing"
	"time"
)

func TestGetAccessToken(t *testing.T) {
	app := NewApp(
		defaultBaseUrl,
		"ww089393e8dd4c985e",
		1000002,
		"Hm9Zq2Pwe8NxIOl1izpkYxQIPFMgUqduooQzt4VGZuU",
		"qiHeY2kCXxn7qRjoUtQcHcSXinNLPCmDvVKMGKLi38b",
		"FZ8XW8Miacl6O3F6qHMV6z04K9BnAL", time.Second*10,
	)
	token, err := app.GetToken()
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Log(token.Token, token.ExpiresIn)
}
