package util

// RefundCheckDto 结构体
type RefundCheckDto struct {
	// 审核状态 恒为 SUCCESS
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 审核描述
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 审核操作时间
	OperateTime string `json:"operate_time,omitempty" xml:"operate_time,omitempty"`
	// 退款单ID
	RefundId int64 `json:"refund_id,omitempty" xml:"refund_id,omitempty"`
	// 主订单ID
	Tid int64 `json:"tid,omitempty" xml:"tid,omitempty"`
	// 子订单ID
	Oid int64 `json:"oid,omitempty" xml:"oid,omitempty"`
	// 退款金额
	RefundFee int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
}
