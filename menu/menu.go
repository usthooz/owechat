package menu

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/swxctx/ghttp"
	"github.com/usthooz/oozlog/go"
	"github.com/usthooz/owechat/config"
	"github.com/usthooz/owechat/token"
)

// Create 创建菜单
func Create(b []*Button) error {
	menuReq := &Menu{
		Button: b,
		Appid:  cfg.BaseConf.Appid,
	}

	ozlog.Debugf("params: %#v", menuReq)

	accessToken, err := token.GetAccessToken()
	if err != nil {
		return err
	}
	ozlog.Infof("token-> %s", accessToken)
	// new request
	req := ghttp.Request{
		Url:       fmt.Sprintf("%s%s", MENU_CREATE, accessToken),
		Body:      menuReq,
		Method:    "POST",
		ShowDebug: cfg.BaseConf.Debug,
	}

	// send request
	resp, err := req.Do()
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Menu Create err, status -> %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	respBs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	ozlog.Debugf("resp: %s", string(respBs))

	var (
		menuResp *ErrorResp
	)
	err = json.Unmarshal(respBs, &menuResp)
	if err != nil {
		return err
	}
	if menuResp.Errcode != 0 || len(menuResp.Errmsg) > 0 {
		return fmt.Errorf("Menu Create: errcode-> %d, errmsg-> %s", menuResp.Errcode, menuResp.Errmsg)
	}
	return nil
}

// Delete 删除菜单
func Delete() error {
	accessToken, err := token.GetAccessToken()
	if err != nil {
		return err
	}
	// new request
	req := ghttp.Request{
		Url:       fmt.Sprintf("%s%s", MENU_DELETE, accessToken),
		Method:    "GET",
		ShowDebug: cfg.BaseConf.Debug,
	}

	// send request
	resp, err := req.Do()
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Menu Create err, status -> %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	respBs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	ozlog.Debugf("resp: %s", string(respBs))

	var (
		menuResp *ErrorResp
	)
	err = json.Unmarshal(respBs, &menuResp)
	if err != nil {
		return err
	}
	if menuResp.Errcode != 0 || len(menuResp.Errmsg) > 0 {
		return fmt.Errorf("Menu Create: errcode-> %d, errmsg-> %s", menuResp.Errcode, menuResp.Errmsg)
	}
	return nil
}
