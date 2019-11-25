package menu

// Menu 菜单内容
type Menu struct {
	Button []*Button `xml:"button"`
	Appid  string    `xml:"appid,omitempty"`
}

// Button 一级菜单
type Button struct {
	Type      string `xml:"type,omitempty"`
	Name      string `xml:"name,omitempty"`
	Key       string `xml:"key,omitempty"`
	SubButton string `xml:"sub_button,omitempty"`
	Url       string `xml:"url,omitempty"`
	MediaId   string `xml:"media_id,omitempty"`
	Pagepath  string `xml:"pagepath,omitempty"`
}

// SubButton 二级菜单内容
type SubButton struct {
	Type     string `xml:"type,omitempty"`
	Name     string `xml:"name,omitempty"`
	Key      string `xml:"key,omitempty"`
	Url      string `xml:"url,omitempty"`
	MediaId  string `xml:"media_id,omitempty"`
	Pagepath string `xml:"pagepath,omitempty"`
}

// ErrorResp 返回结构
type ErrorResp struct {
	Errcode int32  `xml:"errcode"`
	Errmsg  string `xml:"errmsg"`
}
