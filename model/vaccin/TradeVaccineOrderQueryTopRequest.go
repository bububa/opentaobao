package vaccin

// TradeVaccineOrderQueryTopRequest 结构体
type TradeVaccineOrderQueryTopRequest struct {
	// 卖家ID
	MerchantId string `json:"merchant_id,omitempty" xml:"merchant_id,omitempty"`
	// 业务订单号
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
}
