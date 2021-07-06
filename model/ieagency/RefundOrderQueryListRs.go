package ieagency

// RefundOrderQueryListRs 结构体
type RefundOrderQueryListRs struct {
	// 退票申请单列表
	RefundOrderSimpleVos []RefundOrderSimpleVo `json:"refund_order_simple_vos,omitempty" xml:"refund_order_simple_vos>refund_order_simple_vo,omitempty"`
	// 错误原因
	ApiErrorMsg string `json:"api_error_msg,omitempty" xml:"api_error_msg,omitempty"`
	// 错误原因
	ErrTrace string `json:"err_trace,omitempty" xml:"err_trace,omitempty"`
	// 机器名称
	HostName string `json:"host_name,omitempty" xml:"host_name,omitempty"`
	// API错误码
	ApiErrorCode int64 `json:"api_error_code,omitempty" xml:"api_error_code,omitempty"`
	// failure
	Failure bool `json:"failure,omitempty" xml:"failure,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
