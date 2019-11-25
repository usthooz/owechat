package cfg

import (
	"github.com/usthooz/owechat/cache"
	"github.com/xiaoenai/tp-micro/model/redis"
)

var (
	BaseConf *BaseConfig
)

// BaseConfig owechat基础配置中心
type BaseConfig struct {
	// wx appid
	Appid string `yaml:"appid"`
	// wx app secret
	AppSecret string `yaml:"app_secret"`
	// redis缓存的头部key(选填)
	CacheKey string `yaml:"cache_key"`
	// Debug 是否调试模式
	Debug bool `yaml:"debug"`
}

// Config owechat配置中心
type Config struct {
	// Base 基础配合
	Base *BaseConfig `yaml:"base"`
	// Redis 缓存的redis配置
	Redis *redis.Config `yaml:"redis"`
}

// InitOwechat
func (cfg *Config) Reload() error {
	// init base conf
	BaseConf = cfg.Base
	// init cache redis
	if err := cache.InitCache(cfg.Redis, BaseConf.CacheKey); err != nil {
		return err
	}
	return nil
}
