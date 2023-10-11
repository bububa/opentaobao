package tmallsc

// ServiceInsuranceCallbackRequest 结构体
type ServiceInsuranceCallbackRequest struct {
	// 保单号
	InsuranceOrderNo string `json:"insurance_order_no,omitempty" xml:"insurance_order_no,omitempty"`
	// 投保时间
	InsuranceStartTime string `json:"insurance_start_time,omitempty" xml:"insurance_start_time,omitempty"`
	// 退保时间(暂定）
	InsuranceRefundTime string `json:"insurance_refund_time,omitempty" xml:"insurance_refund_time,omitempty"`
	// 投保数量
	InsuranceCount int64 `json:"insurance_count,omitempty" xml:"insurance_count,omitempty"`
	// 投保对应服务单
	InsuranceServiceOrderId int64 `json:"insurance_service_order_id,omitempty" xml:"insurance_service_order_id,omitempty"`
	// 投保类型 1：只换不修；2：碎屏；3：延长保修；4:全面保修;5:意外保；6:无忧退
	InsuranceType int64 `json:"insurance_type,omitempty" xml:"insurance_type,omitempty"`
	//  保单状态 1 投保成功； 99 投保失败
	InsuranceOrderStatus int64 `json:"insurance_order_status,omitempty" xml:"insurance_order_status,omitempty"`
}
