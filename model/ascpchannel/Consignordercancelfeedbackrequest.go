package ascpchannel

// Consignordercancelfeedbackrequest 结构体
type Consignordercancelfeedbackrequest struct {
	// 供应商id
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 履约单号
	BizOrderCode string `json:"biz_order_code,omitempty" xml:"biz_order_code,omitempty"`
	// 业务时间
	BizTime string `json:"biz_time,omitempty" xml:"biz_time,omitempty"`
	// 取消失败原因
	CancelReason string `json:"cancel_reason,omitempty" xml:"cancel_reason,omitempty"`
	// 一盘货业务模式，默认为0代表商家仓商家配，为1代表商家仓自营配 (为1时会强制校验配CP和单号必须与取号时一致，且多包裹必须一次性发货)
	BusinessModel string `json:"business_model,omitempty" xml:"business_model,omitempty"`
	// 是否取消成功,true成功/false失败
	CancelResult bool `json:"cancel_result,omitempty" xml:"cancel_result,omitempty"`
}
