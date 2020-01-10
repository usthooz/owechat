package cache

import (
	"github.com/xiaoenai/tp-micro/model/redis"
)

// wxRedis 缓存accexx_token的redis
var wxRedis = struct {
	*redis.Client
	module *redis.Module
}{
	module: redis.NewModule("owx"),
}

// InitCache
func InitCache(cfg *redis.Config, cacheKey ...string) error {
	// init redis client
	cli, err := redis.NewClient(cfg)
	if err != nil {
		return err
	}
	wxRedis.Client = cli
	// init cache key
	accessTokenKey = wxRedis.module.Key(defaultTokenCacheKey)
	if len(cacheKey) > 0 && len(cacheKey[0]) > 0 {
		accessTokenKey = cacheKey[0]
	}
	// ticket cache key
	ticketCacheKey = wxRedis.module.Key(defaultTicketCacheKey)
	if len(cacheKey) > 2 && len(cacheKey[1]) > 0 {
		ticketCacheKey = cacheKey[1]
	}
	return nil
}
