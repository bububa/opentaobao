package idle

// PayQueryRequest 结构体
type PayQueryRequest struct {
	// 业务订单号
	BizOrderId string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
}
