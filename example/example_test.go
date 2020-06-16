package example

import (
	"fmt"
	"testing"
	"time"

	"github.com/swxctx/xlog"
	cfg "github.com/usthooz/owechat/config"
	"github.com/usthooz/owechat/jssign"
	"github.com/usthooz/owechat/menu"
	"github.com/usthooz/owechat/msg"
	"github.com/usthooz/owechat/pay"
	"github.com/usthooz/owechat/token"
	"github.com/usthooz/owechat/util"
	"github.com/xiaoenai/tp-micro/model/redis"
)

func init() {
	config := cfg.Config{
		Base: &cfg.BaseConfig{
			// 测试号
			Appid:     "wxda4390c918b1fe8e",
			AppSecret: "e8fba0b3df83a545d4ddae8e7bcbfd59",
			// AesKey:    "rfBd56ti2SMtYvSgD5xAV0YU99zampta7Z7S575KLkI",
			// Token:     "xiaoenai",
			Debug: true,
		},
		Redis: redis.NewConfig(),
	}
	if err := config.Reload(); err != nil {
		xlog.Errorf("init err-> %v", err)
	}
}

// TestGetToken
func TestGetToken(t *testing.T) {
	token, err := token.GetAccessToken()
	if err != nil {
		xlog.Errorf("GetTokenTest err-> %v", err)
	}
	xlog.Infof("Token-> %s", token)
}

// TestCreateMenu 创建菜单
func TestCreateMenu(t *testing.T) {
	var (
		buttons []*menu.Button
	)
	buttons = append(buttons, &menu.Button{
		Type: menu.MenuByClick,
		Name: "领取红包",
		Key:  "XEA_GARDEN_GET_PACK",
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
		Signature:    "9af52b2de12ba0a97196822b6b17dece94bc1d7e",
		Timestamp:    1574938655,
		Nonce:        "Yl9HxD4QWS96igWF",
		MsgSignature: "afec8c5cb589c7d0924c296eb732819e166cc4a6",
		EncryptType:  "aes",
		EncryptedMsg: "2uLH+3HnAoPmzbbqWbEL8yHD1NLDxmQ0b4xnDd6dNi3LYyDg44D8VJfscbuR+2zrEZwEq7rU1eEzkOizllAMjyu7s4SOqRbB8PCDmKnHe+2e8/6LgBFj4OcF9t6K0xw7N4yiUYi/ZtThIv0md6ad6ecm4R+dLpwtYbWR7RRA/5PbB7NN1kzVzavdlrSryj6YC4VjWsfljjf8OIYyj6+XwpEaOe7cXidRi0hHwLpfrdiOsEWM+QZKCG7ziOlL2zm8D5lMdUrdKUwrt80GNXbOg22LIDMHSzVpztiTHOS38uGcwdNdpS4O1p08K7MIMXIttxA084/CB8Oe657EsC5k1g==",
	})
	if err != nil {
		t.Errorf("Err-> %v", err)
	}
	t.Logf("Msg-> %#v", msg)
}

// TestSign
func TestSign(t *testing.T) {
	// sign := util.Signature("xiaoenai", "1574934213", "1251643255")
	sign := util.Signature("xiaoenai", "1574934213", "1251643255", "YXPHWrA3a9eRyjubuU0xjdtxWdx9duHkncV/iEk6Qv1KChAbG5PO471St0D+CWy/l0PQCWioR05eg6nCD+JaJuJpTgqnqtfdFuRKBopWVDDaYmRKxNj6w7PQzEpsEYZ6FfT91pYOJCjZCasbpfhnJMXQlqTM8zBVkyxzSDjyg7ul6OZeKIDjI9SiRwai7rTOGLzX2aGSCLeE2z6AhTNomYePTT0uE9VL418cNWdAu9lqVlR1Te8b6Dxxt1BnZGL3cneAfk1fxNlIbraNQndC0UvA7hXQyo3JNbqhwY1eiHKZDsQCY/0QKlbCMRN6Eck1Sb8wEU2WMYN367H/397o2Mz0T7EA10H4kd7BRi+eQVx9i4qKtbGub6YRu7tqUnObHUYETSVjYb2r6Fc48qZcIBSHSuKufPighTLdAQbslCs=")
	t.Logf("Sign-> %v", sign)
}

// TestSendPack
func TestSendPack(t *testing.T) {
	mchBillno := fmt.Sprintf("%d", time.Now().Unix())
	resp, err := pay.SendPack(&pay.SendredPackParams{
		MchBillno:   mchBillno,
		SendName:    "ooz",
		ReOpenid:    "oxQ-Tt47z8svNzqFXjKCABvhTCxU",
		TotalAmount: 10,
		TotalNum:    1,
		Wishing:     "恭喜发财",
		ActName:     "一起拼",
		Remark:      mchBillno,
		SceneId:     pay.SCENEBYPRODUCT_1,
	})
	if err != nil {
		t.Errorf("Err-> %v", err)
	}
	t.Logf("resp-> %#v", resp)
}

// TestJsSign
func TestJsSign(t *testing.T) {
	url, err := jssign.GetJsSign("https://www.ooz.ink")
	if err != nil {
		t.Errorf("Err-> %v", err)
	}
	t.Logf("Sign Url-> %s", url)
}
