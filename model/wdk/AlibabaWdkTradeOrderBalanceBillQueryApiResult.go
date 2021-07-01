package wdk

// AlibabaWdkTradeOrderBalanceBillQueryApiResult 结构体
type AlibabaWdkTradeOrderBalanceBillQueryApiResult struct {
	// model
	Model *OrderBalanceBillResponseDo `json:"model,omitempty" xml:"model,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 成功失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
