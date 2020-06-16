package jssign

import (
	"crypto/sha1"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/usthooz/gutil"
	cfg "github.com/usthooz/owechat/config"
	"github.com/usthooz/owechat/ticket"
)

// GetJsSign
func GetJsSign(url string) (*JsSign, error) {
	jsTicket, err := ticket.GetTicket()
	if err != nil {
		return nil, err
	}
	// splite url
	urlSlice := strings.Split(url, "#")
	jsSign := &JsSign{
		Appid:     cfg.BaseConf.Appid,
		Noncestr:  gutil.RandString(16),
		Timestamp: strconv.FormatInt(time.Now().UTC().Unix(), 10),
		Url:       urlSlice[0],
	}
	jsSign.Signature = signature(jsTicket, jsSign.Noncestr, jsSign.Timestamp, jsSign.Url)
	return jsSign, nil
}

// signature
func signature(jsTicket, noncestr, timestamp, url string) string {
	h := sha1.New()
	h.Write([]byte(fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%s&url=%s", jsTicket, noncestr, timestamp, url)))
	return fmt.Sprintf("%x", h.Sum(nil))
}
