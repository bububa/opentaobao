package alicom

// DistributeTradeOrderInfo 结构体
type DistributeTradeOrderInfo struct {
	// 天猫订单号
	TbTradeId string `json:"tb_trade_id,omitempty" xml:"tb_trade_id,omitempty"`
	// 交易请求流水号
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 金额(单位:分)
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 支付跳转地址
	PayUrl string `json:"pay_url,omitempty" xml:"pay_url,omitempty"`
	// 加密串
	SignStr string `json:"sign_str,omitempty" xml:"sign_str,omitempty"`
	// 0:初始化，1:成功，2:失败
	Status string `json:"status,omitempty" xml:"status,omitempty"`
}
