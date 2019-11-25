package msg

type MessageRecive struct {
	// 开发者微信号
	ToUserName string `xml:"ToUserName,omitempty"`
	// 发送方帐号（一个OpenID）
	FromUserName string `xml:"FromUserName,omitempty"`
	// 消息创建时间 （整型）
	CreateTime int64 `xml:"CreateTime,omitempty"`
	// 消息类型
	MsgType string `xml:"MsgType,omitempty"`
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

type MessageReply struct {
	// 开发者微信号
	ToUserName string `xml:"ToUserName,omitempty"`
	// 发送方帐号（一个OpenID）
	FromUserName string `xml:"FromUserName,omitempty"`
	// 消息创建时间 （整型）
	CreateTime int64 `xml:"CreateTime,omitempty"`
	// 消息类型
	MsgType string `xml:"MsgType,omitempty"`
	// 文本消息内容
	Content string `xml:"Content,omitempty"`
	// 通过素材管理中的接口上传多媒体文件，得到的id
	MediaId string `xml:"url,omitempty"`
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
	// 图文消息个数；当用户发送文本、图片、视频、图文、地理位置这五种消息时，开发者只能回复1条图文消息；其余场景最多可回复8条图文消息
	ArticleCount int32 `xml:"ArticleCount,omitempty"`
	// 图文消息信息，注意，如果图文数超过限制，则将只发限制内的条数
	Articles string `xml:"Articles,omitempty"`
	// 图片链接，支持JPG、PNG格式，较好的效果为大图360*200，小图200*200
	PicUrl string `xml:"PicUrl,omitempty"`
	// 点击图文消息跳转链接
	Url string `xml:"Url,omitempty"`
}
