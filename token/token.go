package token

import (
	"fmt"
	"time"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/usthooz/owechat/cache"
	"github.com/usthooz/owechat/config"
	"github.com/usthooz/owechat/util"
)

// GetAccessToken 获取普通调用access_token ［有次数限制，必须缓存］
func GetAccessToken() (accessToken string, err error) {
	// 从redis取出token
	accessToken, err = cache.GetAccessToken()
	if err != nil {
		return
	}
	if len(accessToken) > 0 {
		return accessToken, nil
	}
	/*
		{
			"access_token":"bxLdikRXVbTPdHSM05e5u5sUoXNKd8-41ZO3MhKoyN5OfkWITDGgnr2fwJ0m9E8NYzWKVZvdVtaUgWvsdshFKA",
			"expires_in":7200
		}
	*/

	url := fmt.Sprintf("%s%sappid=%s&secret=%s", API_URL_PREFIX, AUTH_URL, cfg.BaseConf.Appid, cfg.BaseConf.AppSecret)

	// request
	_, bs, e := util.GET(url)
	if e != nil {
		err = fmt.Errorf("请求失败, err=%v", e)
		return
	}

	var (
		j *simplejson.Json
	)
	j, e = simplejson.NewJson(bs)
	if e != nil {
		err = fmt.Errorf("解析失败, err=%v", err)
		return
	}

	// 判断获取是否存在错误
	if _, ok := j.CheckGet("errcode"); ok {
		//获取失败
		err = fmt.Errorf("获取access_token失败, err=%s", string(bs))
		return
	}

	// 获取ak并缓存
	accessToken = j.GetPath("access_token").MustString()

	// 缓存有效期
	expire := j.GetPath("expires_in").MustInt()
	if expire < 1 {
		expire = 3600
	} else {
		//留个缓冲
		expire = expire - 200
	}
	// 缓存到redis
	err = cache.SetAccessToken(accessToken, time.Duration(expire)*time.Second)
	return
}
