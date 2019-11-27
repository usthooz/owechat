package cache

import (
	"time"

	"github.com/xiaoenai/tp-micro/model/redis"
)

// SetAccessToken 设置access token
func SetAccessToken(token string, duration time.Duration) error {
	return wxRedis.Set(accessTokenKey, token, duration).Err()
}

// GetAccessToken 获取accesstoken
func GetAccessToken() (string, error) {
	res := wxRedis.Get(accessTokenKey)
	if res.Err() == redis.Nil {
		return "", nil
	}
	return res.Val(), res.Err()
}
