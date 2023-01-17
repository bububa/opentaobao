package idle

// ConsignmentOrderQuery 结构体
type ConsignmentOrderQuery struct {
	// appkey
	AppKey string `json:"app_key,omitempty" xml:"app_key,omitempty"`
	// 来源
	Source string `json:"source,omitempty" xml:"source,omitempty"`
	// 订单id
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 服务商id
	BuyerId int64 `json:"buyer_id,omitempty" xml:"buyer_id,omitempty"`
}
