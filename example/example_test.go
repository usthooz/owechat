package example

import (
	"testing"

	"github.com/usthooz/oozlog/go"
	"github.com/usthooz/owechat/config"
	"github.com/usthooz/owechat/menu"
	"github.com/usthooz/owechat/msg"
	"github.com/usthooz/owechat/token"
	"github.com/xiaoenai/tp-micro/model/redis"
)

func init() {
	config := cfg.Config{
		Base: &cfg.BaseConfig{
			Appid:     "wxda4390c918b1fe8e",
			AppSecret: "e8fba0b3df83a545d4ddae8e7bcbfd59",
			AesKey:    "rfBd56ti2SMtYvSgD5xAV0YU99zampta7Z7S575KLkI",
			Token:     "xiaoenai",
			Debug:     true,
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

// TestCreateMenu 创建菜单
func TestCreateMenu(t *testing.T) {
	var (
		buttons []*menu.Button
	)
	buttons = append(buttons, &menu.Button{
		Type: menu.MenuByClick,
		Name: "点我点我",
		Key:  "V1001_TODAY_MUSIC",
	})
	err := menu.Create(buttons)
	if err != nil {
		t.Errorf("Err-> %v", err)
	}
}

// TestCreateMenu 床架菜单
func TestDeleteMenu(t *testing.T) {
	err := menu.Delete()
	if err != nil {
		t.Errorf("Err-> %v", err)
	}
}

// TestDecryptMsg
func TestDecryptMsg(t *testing.T) {
	msg, err := msg.DecodeMsg(&msg.EncryptedXMLMsg{
		// Signature:    "ccfe3a24491fb76be939a32d0e50bc3ecb7a9ca5",
		// Timestamp:    1574867028,
		// Nonce:        "1761970467",
		EncryptType:  "aes",
		MsgSignature: "5066d943c76d443ed86a4b8e68013fb00a62cb95",
	})
	if err != nil {
		panic(err)
	}
	t.Logf("Msg-> %#v", msg)
}
