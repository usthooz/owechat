package ticket

import (
	"fmt"
	"time"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/usthooz/owechat/cache"
	"github.com/usthooz/owechat/token"
	"github.com/usthooz/owechat/util"
)

// GetTicket
func GetTicket() (ticket string, err error) {
	// 首先从缓存获取
	ticket, err = cache.GetTicket()
	if len(ticket) > 0 {
		return ticket, nil
	}

	// 重新获取ticket，获取token
	accessToken, e := token.GetAccessToken()
	if e != nil {
		err = e
		return
	}

	// req url
	url := fmt.Sprintf("%s%saccess_token=%s&type=jsapi", WxAPIURLPrefix, WxGetTicketURL, accessToken)

	_, bs, e := util.GET(url)
	if e != nil {
		err = fmt.Errorf("GetTicket: get ticket err-> %v", e)
		return
	}

	var (
		resp *simplejson.Json
	)
	resp, e = simplejson.NewJson(bs)
	if e != nil {
		err = fmt.Errorf("GetTicket: Unmarshal err-> %v", err)
		return
	}

	/*
		Response:
		{
			"errcode":0,
			"errmsg":"ok",
			"ticket":"bxLdikRXVbTPdHSM05e5u5sUoXNKd8",
			"expires_in":7200
		}
	*/
	if _, ok := resp.CheckGet("ticket"); !ok {
		err = fmt.Errorf("GetTicket: get ticket err-> %s", string(bs))
		return
	}

	// get ticket
	ticket = resp.GetPath("ticket").MustString()
	if len(ticket) == 0 {
		err = fmt.Errorf("GetTicket: ticket is nil.")
		return
	}

	expire := resp.GetPath("expires_in").MustInt()
	if expire < 1 {
		expire = 3600
	} else {
		expire = expire - 100
	}
	err = cache.SetTicket(ticket, time.Duration(expire)*time.Second)

	return
}
