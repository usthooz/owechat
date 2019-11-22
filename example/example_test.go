package example

import (
	"testing"

	"github.com/usthooz/oozlog/go"
	"github.com/usthooz/owechat/config"
	"github.com/usthooz/owechat/token"
	"github.com/xiaoenai/tp-micro/model/redis"
)

func init() {
	config := cfg.Config{
		Base: &cfg.BaseConfig{
			Appid:     "wx8e452f637b9280af",
			AppSecret: "c1975f7a763fac2b4969a512c6701edb",
		},
		Redis: redis.NewConfig(),
	}
	if err := config.Reload(); err != nil {
		ozlog.Errorf("init err-> %v", err)
	}
}

// TestGetToken
func TestGetToken(t *testing.T) {
	token, err := token.GetAccessToken()
	if err != nil {
		ozlog.Errorf("GetTokenTest err-> %v", err)
	}
	ozlog.Infof("Token-> %s", token)
}
