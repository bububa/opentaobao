package jst

// RefundTrace 结构体
type RefundTrace struct {
	// 回流的订单状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 动作发生的时间
	ActionTime string `json:"action_time,omitempty" xml:"action_time,omitempty"`
	// 备注字段
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 应用标题
	AppTitle string `json:"app_title,omitempty" xml:"app_title,omitempty"`
	// 交易tid
	Tid int64 `json:"tid,omitempty" xml:"tid,omitempty"`
	// 退款编号
	RefundId int64 `json:"refund_id,omitempty" xml:"refund_id,omitempty"`
}
