package vaccin

// TradeSubscribeDetailQueryTopRequest 结构体
type TradeSubscribeDetailQueryTopRequest struct {
	// 商家ID
	MerchantId string `json:"merchant_id,omitempty" xml:"merchant_id,omitempty"`
	// 业务订单号
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 预约单主键
	SubscribeId int64 `json:"subscribe_id,omitempty" xml:"subscribe_id,omitempty"`
}
