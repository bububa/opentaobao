package wdk

// FulfillLogisticDefaultResult 结构体
type FulfillLogisticDefaultResult struct {
	// true 调用成功 false 调用失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 返回码含义描述
	ErrDesc string `json:"err_desc,omitempty" xml:"err_desc,omitempty"`
	// 返回码(success=false时不能为空)： SYSTEM_ERROR :系统异常(指令可重发) PARAM_ERROR :参数错误(指令不可重发，监控报警) BUSINESS_ERROR:业务异常(指令不可重发，监控报警)
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
}
