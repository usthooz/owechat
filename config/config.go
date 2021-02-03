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
	// AesKey 消息加密的aes算法key
	AesKey string `yaml:"aes_key"`
	// Token 微信后台配置的token
	Token string `yml:"token"`
	// PayKey 支付秘钥
	PayKey string `yaml:"pay_key"`
	// MchId 商户ID
	MchId string `yaml:"mch_id"`
	// PayCertFile 红包发放证书
	PayCertFile string `yaml:"pay_cert_file"`
	// PayKeyFile 红包发放证书
	PayKeyFile string `yaml:"pay_key_file"`
	// redis缓存的头部key(选填)
	CacheKey string `yaml:"cache_key"`
	// Debug 是否调试模式
	Debug bool `yaml:"debug"`
	// TimeoutSecond 超时时间
	TimeoutSecond int64 `yaml:"timeout_second"`
}

// Config owechat配置中心
type Config struct {
	// Base 基础配合
	Base *BaseConfig `yaml:"base"`
	// Redis 缓存的redis配置
	Redis *redis.Config `yaml:"redis"`
}

// NewConfig
func NewConfig() *Config {
	return &Config{
		Base: &BaseConfig{
			Debug: true,
		},
		Redis: redis.NewConfig(),
	}
}

// InitOwechat
func (cfg *Config) Reload() error {
	// init base conf
	BaseConf = cfg.Base
	if BaseConf.TimeoutSecond <= 0 {
		// 默认2s超时
		BaseConf.TimeoutSecond = 2
	}
	// init cache redis
	if err := cache.InitCache(cfg.Redis, BaseConf.CacheKey); err != nil {
		return err
	}
	return nil
}
