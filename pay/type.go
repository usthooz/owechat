package pay

/*
	https://pay.weixin.qq.com/wiki/doc/api/tools/cash_coupon.php?chapter=13_4&index=3
*/

// SendredPackReq 发放现金红包请求参数
type SendredPackReq struct {
	// 随机字符串，不长于32位
	NonceStr string `xml:"nonce_str,omitempty"`
	// 请求签名
	Sign string `xml:"sign,omitempty"`
	// 商户订单号（每个订单号必须唯一。取值范围：0~9，a~z，A~Z）
	MchBillno string `xml:"mch_billno,omitempty"`
	// 商户号
	MchId string `xml:"mch_id,omitempty"`
	// 微信分配的公众账号ID（企业号corpid即为此appId）。在微信开放平台（open.weixin.qq.com）申请的移动应用appid无法使用该接口。
	Wxappid string `xml:"wxappid,omitempty"`
	// 红包发送者名称 注意：敏感词会被转义成字符*
	SendName string `xml:"send_name,omitempty"`
	// 接受红包的用户openid
	ReOpenid string `xml:"re_openid,omitempty"`
	// 付款金额，单位分
	TotalAmount int32 `xml:"total_amount,omitempty"`
	// 红包发放总人数 total_num=1
	TotalNum int32 `xml:"total_num,omitempty"`
	// 红包祝福语 注意：敏感词会被转义成字符*
	Wishing string `xml:"wishing,omitempty"`
	// 调用接口的机器Ip地址
	ClientIp string `xml:"client_ip,omitempty"`
	// 活动名称 注意：敏感词会被转义成字符*
	ActName string `xml:"act_name,omitempty"`
	// 猜越多得越多，快来抢！
	Remark string `xml:"remark,omitempty"`
	/*
		发放红包使用场景，红包金额大于200或者小于1元时必传
		PRODUCT_1:商品促销
		PRODUCT_2:抽奖
		PRODUCT_3:虚拟物品兑奖
		PRODUCT_4:企业内部福利
		PRODUCT_5:渠道分润
		PRODUCT_6:保险回馈
		PRODUCT_7:彩票派奖
		PRODUCT_8:税务刮奖
	*/
	SceneId string `xml:"scene_id,omitempty"`
	// 活动信息
	RiskInfo string `xml:"risk_info,omitempty"`
}

// SendredPackParams 发放现金红包请求参数
type SendredPackParams struct {
	// 商户订单号（每个订单号必须唯一。取值范围：0~9，a~z，A~Z）
	MchBillno string `xml:"mch_billno,omitempty"`
	// 红包发送者名称 注意：敏感词会被转义成字符*
	SendName string `xml:"send_name,omitempty"`
	// 接受红包的用户openid
	ReOpenid string `xml:"re_openid,omitempty"`
	// 付款金额，单位分
	TotalAmount int32 `xml:"total_amount,omitempty"`
	// 红包发放总人数 total_num=1
	TotalNum int32 `xml:"total_num,omitempty"`
	// 红包祝福语 注意：敏感词会被转义成字符*
	Wishing string `xml:"wishing,omitempty"`
	// 活动名称 注意：敏感词会被转义成字符*
	ActName string `xml:"act_name,omitempty"`
	// 猜越多得越多，快来抢！
	Remark string `xml:"remark,omitempty"`
	/*
		发放红包使用场景，红包金额大于200或者小于1元时必传
		PRODUCT_1:商品促销
		PRODUCT_2:抽奖
		PRODUCT_3:虚拟物品兑奖
		PRODUCT_4:企业内部福利
		PRODUCT_5:渠道分润
		PRODUCT_6:保险回馈
		PRODUCT_7:彩票派奖
		PRODUCT_8:税务刮奖
	*/
	SceneId string `xml:"scene_id,omitempty"`
	// 活动信息
	RiskInfo *RiskInfo `xml:"risk_info,omitempty"`
}

// RiskInfo 活动信息
type RiskInfo struct {
	// 用户操作的时间戳
	Posttime int64 `xml:"posttime,omitempty"`
	// 业务系统账号的手机号，国家代码-手机号。不需要+号
	Mobile string `xml:"mobile,omitempty"`
	// mac 地址或者设备唯一标识
	Deviceid string `xml:"deviceid,omitempty"`
	// 用户操作的客户端版本 把值为非空的信息用key=value进行拼接，再进行urlencode urlencode(posttime=xx& mobile =xx&deviceid=xx)
	ClientVersion string `xml:"clientversion,omitempty"`
}

// CommonPackResp 公用返回结构
type CommonPackResp struct {
	// SUCCESS/FAIL 此字段是通信标识，非红包发放结果标识，红包发放是否成功需要查看result_code来判断
	ReturnCode string `xml:"return_code"`
	// 返回信息，如非空，为错误原因 签名失败 参数格式校验错误
	ReturnMsg string `xml:"return_msg"`
	// SUCCESS/FAIL 注意：当状态为FAIL时，存在业务结果未明确的情况。所以如果状态是FAIL，请务必再请求一次查询接口[请务必关注错误代码（err_code字段），通过查询得到的红包状态确认此次发放的结果。]，以确认此次发放的结果。
	ResultCode string `xml:"result_code"`
	// 错误码信息 注意：出现未明确的错误码（SYSTEMERROR等）时，请务必用原商户订单号重试，或者再请求一次查询接口以确认此次发放的
	ErrCode string `xml:"err_code"`
	// 结果信息描述
	ErrCodeDes string `xml:"err_code_des"`
	// 商户订单号（每个订单号必须唯一） 组成：mch_id+yyyymmdd+10位一天内不能重复的数字
	MchBillno string `xml:"mch_billno"`
	// 微信支付分配的商户号
	MchId string `xml:"mch_id"`
}

// SendredPackResp 现金红包发放响应结构
type SendredPackResp struct {
	CommonPackResp
	// 微信分配的公众账号ID（企业号corpid即为此appId）。在微信开放平台（open.weixin.qq.com）申请的移动应用appid无法使用该接口。
	Wxappid string `xml:"wxappid,omitempty"`
	// 接受红包的用户openid
	ReOpenid string `xml:"re_openid,omitempty"`
	// 付款金额，单位分
	TotalAmount int32 `xml:"total_amount,omitempty"`
	// 红包订单的微信单号
	SendListid string `xml:"send_listid,omitempty"`
}

/*
	https://pay.weixin.qq.com/wiki/doc/api/tools/cash_coupon.php?chapter=13_6&index=5
*/
// QueryPackReq 查询红包记录参数
type QueryPackReq struct {
	// 随机字符串，不长于32位
	NonceStr string `xml:"nonce_str,omitempty"`
	// 请求签名
	Sign string `xml:"sign,omitempty"`
	// 商户订单号（每个订单号必须唯一。取值范围：0~9，a~z，A~Z）
	MchBillno string `xml:"mch_billno,omitempty"`
	// 商户号
	MchId string `xml:"mch_id,omitempty"`
	// 微信分配的公众账号ID（企业号corpid即为此appId）。在微信开放平台（open.weixin.qq.com）申请的移动应用appid无法使用该接口。
	Appid string `xml:"appid,omitempty"`
	// MCHT:通过商户订单号获取红包信息。
	BillType string `xml:"bill_type,omitempty"`
}

// QueryPackResp 查询红包记录响应
type QueryPackResp struct {
	CommonPackResp
	// 使用API发放现金红包时返回的红包单号
	DetailId string `xml:"detail_id,omitempty"`
	// SENDING:发放中 SENT:已发放待领取 FAILED：发放失败 RECEIVED:已领取 RFUND_ING:退款中 REFUND:已退款
	Status string `xml:"status,omitempty"`
	// API:通过API接口发放 UPLOAD:通过上传文件方式发放 ACTIVITY:通过活动方式发放
	SendType string `xml:"send_type,omitempty"`
	// GROUP:裂变红包 NORMAL:普通红包
	HbType string `xml:"hb_type,omitempty"`
	// 红包个数
	TotalNum int32 `xml:"total_num,omitempty"`
	// 红包总金额（单位分）
	TotalAmount int32 `xml:"total_amount,omitempty"`
	// 发送失败原因
	Reason string `xml:"reason,omitempty"`
	// 发送时间
	SendTime string `xml:"send_time,omitempty"`
	// 退款时间
	RefundTime string `xml:"refund_time,omitempty"`
	// 退款金额
	RefundAmount int32 `xml:"refund_amount,omitempty"`
	// 祝福语
	Wishing string `xml:"wishing,omitempty"`
	// 活动描述
	Remark string `xml:"remark,omitempty"`
	// 活动名称
	ActName string `xml:"act_name,omitempty"`
	// 裂变红包的领取列表
	Hblist []*HbInfo `xml:"hblist,omitempty"`
}

// HbInfo 裂变红包信息
type HbInfo struct {
	// 领取红包的openid
	Openid string `xml:"openid,omitempty"`
	// 领取金额
	Amount int32 `xml:"amount,omitempty"`
	// 领取红包的时间 2015-04-21 20:00:00
	RcvTime string `xml:"rcv_time,omitempty"`
}
