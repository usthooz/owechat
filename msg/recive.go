package msg

// CommonMsg 消息通用结构体
type CommonMsg struct {
	// 签名
	Signature string `param:"<query>" xml:"signature"`
	// 秒级时间戳
	Timestamp int64 `param:"<query>" xml:"timestamp"`
	// 随机值
	Nonce string `param:"<query>" xml:"nonce"`
	// 加密方式
	EncryptType string `params:"<query>" xml:"encrypt_type"`
	// 发送者openID
	Openid string `param:"<query>" xml:"openid"`

	// 开发者微信号
	ToUserName string `xml:"ToUserName,omitempty"`
	// 发送方帐号（一个OpenID）
	FromUserName string `xml:"FromUserName,omitempty"`
	// 消息创建时间 （整型）
	CreateTime int64 `xml:"CreateTime,omitempty"`
	// 消息类型
	MsgType MsgType `xml:"MsgType,omitempty"`
}

// MessageRecive 接收微信发送消息的结构体
type MessageRecive struct {
	CommonMsg
	// 文本消息内容
	Content string `xml:"Content,omitempty"`

	// 消息id，64位整型
	MsgId int64 `xml:"MsgId,omitempty"`

	// 图片链接（由系统生成）
	PicUrl string `xml:"PicUrl,omitempty"`
	// 图片消息媒体id，可以调用获取临时素材接口拉取数据。
	MediaId int64 `xml:"MediaId,omitempty"`

	// 语音格式，如amr，speex等
	Format string `xml:"Format,omitempty"`
	// 语音识别结果，UTF8编码
	Recognition string `xml:"Recognition,omitempty"`

	// 视频消息缩略图的媒体id，可以调用多媒体文件下载接口拉取数据。
	ThumbMediaId string `xml:"ThumbMediaId,omitempty"`

	// 地理位置维度
	Location_X float64 `xml:"Location_X,omitempty"`
	// 地理位置经度
	Location_Y float64 `xml:"Location_Y,omitempty"`
	// 地图缩放大小
	Scale int32 `xml:"Scale,omitempty"`
	// 地理位置信息
	Label string `xml:"Label,omitempty"`

	// 消息标题
	Title string `xml:"Title,omitempty"`
	// 消息描述
	Description string `xml:"Description,omitempty"`
	// 消息链接
	Url string `xml:"Url,omitempty"`

	// 事件类型，subscribe(订阅)、unsubscribe(取消订阅)
	Event string `xml:"Event,omitempty"`
	// 事件KEY值，qrscene_为前缀，后面为二维码的参数值
	EventKey string `xml:"EventKey,omitempty"`

	// 二维码的ticket，可用来换取二维码图片
	Ticket string `xml:"Ticket,omitempty"`

	// 地理位置纬度
	Latitude float64 `xml:"Latitude,omitempty"`
	// 地理位置经度
	Longitude float64 `xml:"Longitude,omitempty"`
	// 地理位置精度
	Precision float64 `xml:"Precision,omitempty"`
}

// EncryptedXMLMsg 安全模式下的消息体
type EncryptedXMLMsg struct {
	XMLName struct{} `xml:"xml" json:"-"`

	// 签名
	Signature string `param:"<query>" xml:"signature"`
	// 秒级时间戳
	Timestamp int64 `param:"<query>" xml:"timestamp"`
	// 随机值
	Nonce string `param:"<query>" xml:"nonce"`
	// 发送者openID
	Openid string `param:"<query>" xml:"openid"`
	// 加密方式
	EncryptType string `param:"<query>" xml:"encrypt_type"`
	// 加密
	MsgSignature string `param:"<query>" xml:"msg_signature"`

	// 微信号
	ToUserName string `xml:"ToUserName"`
	// 加密消息体
	EncryptedMsg string `xml:"Encrypt"`
}
