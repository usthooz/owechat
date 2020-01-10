package oauth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/swxctx/ghttp"
	"github.com/usthooz/oozlog/go"
	"github.com/usthooz/owechat/config"
)

// GetAccessToken 获取用户access_token
func GetAccessToken(code string) (*GetAccessTokenResp, error) {
	reqParams := &GetAccessTokenReq{
		Appid:     cfg.BaseConf.Appid,
		Secret:    cfg.BaseConf.AppSecret,
		Code:      code,
		GrantType: "authorization_code",
	}

	// new request
	req := ghttp.Request{
		Url:       OAUTH2_ACCESS_TOKEN_API,
		Method:    "GET",
		Query:     reqParams,
		ShowDebug: cfg.BaseConf.Debug,
	}

	// send request
	resp, err := req.Do()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GetAccessToken err, status -> %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	respBs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ozlog.Debugf("resp: %s", string(respBs))

	var (
		accessTokenResp *GetAccessTokenResp
	)
	err = json.Unmarshal(respBs, &accessTokenResp)
	if err != nil {
		return nil, err
	}
	return accessTokenResp, nil
}
