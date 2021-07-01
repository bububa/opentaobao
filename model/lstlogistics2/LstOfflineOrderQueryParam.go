package lstlogistics2

// LstOfflineOrderQueryParam 结构体
type LstOfflineOrderQueryParam struct {
	// 外部主订单号
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 买家手机号
	BuyerMobile string `json:"buyer_mobile,omitempty" xml:"buyer_mobile,omitempty"`
}
