package msg

import (
	"encoding/xml"
)

// ReplyMsg 消息回复
type ReplyMsg struct {
	// 消息结构体
	MsgData interface{}
}

// ReplyEncryptedMsg 需要返回的消息体
type ReplyEncryptedMsg struct {
	XMLName xml.Name `xml:"xml"`
	// 加密消息体
	EncryptedMsg string `xml:"Encrypt" json:"Encrypt"`
	// 消息签名
	MsgSignature string `xml:"MsgSignature" json:"MsgSignature"`
	// 秒级时间戳
	Timestamp int64 `xml:"TimeStamp" json:"TimeStamp"`
	// 随机值
	Nonce string `xml:"Nonce" json:"Nonce"`
}

// Reply 公用返回结构体
type Reply struct {
	// 开发者微信号
	ToUserName string `xml:"ToUserName,omitempty"`
	// 发送方帐号（一个OpenID）
	FromUserName string `xml:"FromUserName,omitempty"`
	// 消息创建时间 （整型）
	CreateTime int64 `xml:"CreateTime,omitempty"`
	// 消息类型
	MsgType MsgType `xml:"MsgType,omitempty"`
}

// TxtMsg 文本消息
type TxtMsg struct {
	XMLName xml.Name `xml:"xml"`
	// 开发者微信号
	ToUserName string `xml:"ToUserName,omitempty"`
	// 发送方帐号（一个OpenID）
	FromUserName string `xml:"FromUserName,omitempty"`
	// 消息创建时间 （整型）
	CreateTime int64 `xml:"CreateTime,omitempty"`
	// 消息类型
	MsgType MsgType `xml:"MsgType,omitempty"`

	// 文本消息内容
	Content string `xml:"Content,omitempty"`
}

// PicMsg 图片消息
type PicMsg struct {
	XMLName xml.Name `xml:"xml"`
	// 开发者微信号
	ToUserName string `xml:"ToUserName,omitempty"`
	// 发送方帐号（一个OpenID）
	FromUserName string `xml:"FromUserName,omitempty"`
	// 消息创建时间 （整型）
	CreateTime int64 `xml:"CreateTime,omitempty"`
	// 消息类型
	MsgType MsgType `xml:"MsgType,omitempty"`

	// 通过素材管理中的接口上传多媒体文件，得到的id
	MediaId string `xml:"url,omitempty"`
}

// VoiceMsg 语音消息
type VoiceMsg struct {
	XMLName xml.Name `xml:"xml"`
	// 开发者微信号
	ToUserName string `xml:"ToUserName,omitempty"`
	// 发送方帐号（一个OpenID）
	FromUserName string `xml:"FromUserName,omitempty"`
	// 消息创建时间 （整型）
	CreateTime int64 `xml:"CreateTime,omitempty"`
	// 消息类型
	MsgType MsgType `xml:"MsgType,omitempty"`

	// 通过素材管理中的接口上传多媒体文件，得到的id
	MediaId string `xml:"url,omitempty"`
}

// VideoMsg 回复视频消息
type VideoMsg struct {
	XMLName xml.Name `xml:"xml"`
	// 开发者微信号
	ToUserName string `xml:"ToUserName,omitempty"`
	// 发送方帐号（一个OpenID）
	FromUserName string `xml:"FromUserName,omitempty"`
	// 消息创建时间 （整型）
	CreateTime int64 `xml:"CreateTime,omitempty"`
	// 消息类型
	MsgType MsgType `xml:"MsgType,omitempty"`

	// 通过素材管理中的接口上传多媒体文件，得到的id
	MediaId string `xml:"url,omitempty"`
	// 图文/视频/音乐消息的标题
	Title string `xml:"Title,omitempty"`
	// 图文/视频/音乐消息的描述
	Description string `xml:"Description,omitempty"`
}

// MusicMsg 音乐消息
type MusicMsg struct {
	XMLName xml.Name `xml:"xml"`
	// 开发者微信号
	ToUserName string `xml:"ToUserName,omitempty"`
	// 发送方帐号（一个OpenID）
	FromUserName string `xml:"FromUserName,omitempty"`
	// 消息创建时间 （整型）
	CreateTime int64 `xml:"CreateTime,omitempty"`
	// 消息类型
	MsgType MsgType `xml:"MsgType,omitempty"`

	// 图文/视频/音乐消息的标题
	Title string `xml:"Title,omitempty"`
	// 图文/视频/音乐消息的描述
	Description string `xml:"Description,omitempty"`
	// 音乐链接
	MusicURL string `xml:"MusicURL,omitempty"`
	// 高质量音乐链接，WIFI环境优先使用该链接播放音乐
	HQMusicUrl string `xml:"HQMusicUrl,omitempty"`
	// 缩略图的媒体id，通过素材管理中的接口上传多媒体文件，得到的id
	ThumbMediaId string `xml:"ThumbMediaId,omitempty"`
}

// ArticleMsg 图文消息
type ArticleMsg struct {
	XMLName xml.Name `xml:"xml"`
	// 开发者微信号
	ToUserName string `xml:"ToUserName,omitempty"`
	// 发送方帐号（一个OpenID）
	FromUserName string `xml:"FromUserName,omitempty"`
	// 消息创建时间 （整型）
	CreateTime int64 `xml:"CreateTime,omitempty"`
	// 消息类型
	MsgType MsgType `xml:"MsgType,omitempty"`

	// 图文消息个数；当用户发送文本、图片、视频、图文、地理位置这五种消息时，开发者只能回复1条图文消息；其余场景最多可回复8条图文消息
	ArticleCount int32 `xml:"ArticleCount,omitempty"`
	// 图文消息信息，注意，如果图文数超过限制，则将只发限制内的条数
	Articles []*Article `xml:"Articles,omitempty"`
}

// Article 单篇文章
type Article struct {
	Title       string `xml:"Title,omitempty"`
	Description string `xml:"Description,omitempty"`
	PicURL      string `xml:"PicUrl,omitempty"`
	URL         string `xml:"Url,omitempty"`
}
