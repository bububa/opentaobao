package alidoc

// FailedList 结构体
type FailedList struct {
	// 获取处方失败的订单号
	BizOrderId string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 失败信息
	FailedMessage string `json:"failed_message,omitempty" xml:"failed_message,omitempty"`
}
