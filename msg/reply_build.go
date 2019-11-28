package msg

// NewReply 构建回复公用体
func NewReply(toUsername, fromUsername string, createTime int64, msgType MsgType) *Reply {
	return &Reply{
		ToUserName:   toUsername,
		FromUserName: fromUsername,
		CreateTime:   createTime,
		MsgType:      msgType,
	}
}

// NewTxtMsg
func (r *Reply) NewTxtMsg(content string) *TxtMsg {
	return &TxtMsg{
		ToUserName:   r.ToUserName,
		FromUserName: r.FromUserName,
		CreateTime:   r.CreateTime,
		MsgType:      r.MsgType,
		Content:      content,
	}
}
