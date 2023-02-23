package idle

// ServicePlanInfo 结构体
type ServicePlanInfo struct {
	// 售后处理类型，见对接文档
	ServiceType string `json:"service_type,omitempty" xml:"service_type,omitempty"`
	// 反馈信息
	ShowMessage string `json:"show_message,omitempty" xml:"show_message,omitempty"`
	// 仅退款金额(分)
	OnlyRefundCent int64 `json:"only_refund_cent,omitempty" xml:"only_refund_cent,omitempty"`
	// 退货退款金额(分)
	RefundAndReturnCent int64 `json:"refund_and_return_cent,omitempty" xml:"refund_and_return_cent,omitempty"`
}
