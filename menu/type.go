package menu

// Menu 菜单内容
type Menu struct {
	Button []*Button `json:"button"`
	Appid  string    `json:"appid"`
}

// Button 一级菜单
type Button struct {
	Type      string `json:"type,omitempty"`
	Name      string `json:"name,omitempty"`
	Key       string `json:"key,omitempty"`
	SubButton string `json:"sub_button,omitempty"`
	Url       string `json:"url,omitempty"`
	MediaId   string `json:"media_id,omitempty"`
	Pagepath  string `json:"pagepath,omitempty"`
}

// SubButton 二级菜单内容
type SubButton struct {
	Type     string `json:"type,omitempty"`
	Name     string `json:"name,omitempty"`
	Key      string `json:"key,omitempty"`
	Url      string `json:"url,omitempty"`
	MediaId  string `json:"media_id,omitempty"`
	Pagepath string `json:"pagepath,omitempty"`
}

// ErrorResp 返回结构
type ErrorResp struct {
	Errcode int32  `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}
