package trade

// AgreeRefundCheck 结构体
type AgreeRefundCheck struct {
	// 订单id
	DetailOrderId string `json:"detail_order_id,omitempty" xml:"detail_order_id,omitempty"`
	// 流程状态
	DeliveryProcess string `json:"delivery_process,omitempty" xml:"delivery_process,omitempty"`
	// 提示文案
	DeliveryTips string `json:"delivery_tips,omitempty" xml:"delivery_tips,omitempty"`
}
