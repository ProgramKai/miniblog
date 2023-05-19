package domain

type AccessTokenResp struct {
	BasicResp
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}
