package pay

const (
	// OmitInt 加密忽略值
	OmitInt = 0
	// OmitIntString 加密忽略值
	OmitIntString = "0"
)

const (
	// 通过商户订单号获取红包信息。
	PACKTYPEBYMCHT = "MCHT"
)

const (
	// 商品促销
	SCENEBYPRODUCT_1 = "PRODUCT_1"
	// 抽奖
	SCENEBYPRODUCT_2 = "PRODUCT_2"
	// 虚拟物品兑奖
	SCENEBYPRODUCT_3 = "PRODUCT_3"
	// 企业内部福利
	SCENEBYPRODUCT_4 = "PRODUCT_4"
	// 渠道分润
	SCENEBYPRODUCT_5 = "PRODUCT_5"
	// 保险回馈
	SCENEBYPRODUCT_6 = "PRODUCT_6"
	// 彩票派奖
	SCENEBYPRODUCT_7 = "PRODUCT_7"
	// 税务刮奖
	SCENEBYPRODUCT_8 = "PRODUCT_8"
)

const (
	// 发放中
	PACKSTATUSBYSENDING = "SENDING"
	// 已发放待领取
	PACKSTATUSBYSENT = "SENT"
	// 发放失败
	PACKSTATUSBYFAILED = "FAILED"
	// 已领取
	PACKSTATUSBYRECEIVED = "RECEIVED"
	// 退款中
	PACKSTATUSBYRFUND_ING = "RFUND_ING"
	// 已退款
	PACKSTATUSBYREFUND = "REFUND"
)
