package msg

import (
	"encoding/json"
	"encoding/xml"
	"strconv"
	"time"

	"github.com/usthooz/gutil"
	"github.com/usthooz/owechat/config"
	"github.com/usthooz/owechat/util"
)

// DecodeMsg 解密消息
func DecodeMsg(m *EncryptedXMLMsg) (*MessageRecive, error) {
	// 解密消息
	_, msgByte, err := decryptMsg(cfg.BaseConf.Appid, m.MsgSignature, cfg.BaseConf.AesKey)
	if err != nil {
		return nil, err
	}

	var (
		msgRecive *MessageRecive
	)
	// 解析消息体
	err = json.Unmarshal(msgByte, &msgRecive)
	if err != nil {
		return nil, err
	}
	return msgRecive, err
}

// BuildReplyMsg 构建回复消息
func BuildReplyMsg(replyMsg *ReplyMsg) (*ReplyEncryptedMsg, error) {
	msgByte, err := xml.Marshal(replyMsg.MsgData)
	if err != nil {
		return nil, err
	}
	// 对消息进行加密
	encryptedMsg, err := encryptMsg(gutil.RandomBytes(16), msgByte, cfg.BaseConf.Appid, cfg.BaseConf.AesKey)
	if err != nil {
		return nil, err
	}

	timestamp := time.Now().Unix()
	timestampStr := strconv.FormatInt(timestamp, 10)
	// 随机值
	nonce := gutil.RandString(16)
	// 签名
	msgSignature := util.Signature(cfg.BaseConf.Token, timestampStr, nonce, string(encryptedMsg))

	// 回复的消息结构体
	return &ReplyEncryptedMsg{
		EncryptedMsg: string(encryptedMsg),
		MsgSignature: msgSignature,
		Timestamp:    timestamp,
		Nonce:        nonce,
	}, nil
}
