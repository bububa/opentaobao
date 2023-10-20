package wdk

// AlibabaWdkTradeRefundSuccessCreateApiResult 结构体
type AlibabaWdkTradeRefundSuccessCreateApiResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 接口状态
	State bool `json:"state,omitempty" xml:"state,omitempty"`
	// 是否成功 true-成功;false-失败
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
