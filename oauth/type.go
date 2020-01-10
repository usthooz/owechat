package oauth

// GetAccessTokenReq 通过code换取access_token
type GetAccessTokenReq struct {
	Appid  string `json:"appid"`
	Secret string `json:"secret"`
	// 用户授权得到的code
	Code string `json:"code"`
	// 填写为authorization_code
	GrantType string `json:"grant_type"`
}

/*
	access_token	网页授权接口调用凭证,注意：此access_token与基础支持的access_token不同
	expires_in	access_token接口调用凭证超时时间，单位（秒）
	refresh_token	用户刷新access_token
	openid	用户唯一标识
	scope	用户授权的作用域，使用逗号（,）分隔
*/
// GetAccessTokenResp
type GetAccessTokenResp struct {
	AccessToken  string  `json:"access_token"`
	ExpiresIn    float64 `json:"expires_in"`
	RefreshToken string  `json:"refresh_token"`
	Openid       string  `json:"openid"`
	Scope        string  `json:"scope"`
}
