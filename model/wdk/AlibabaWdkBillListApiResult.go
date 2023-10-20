package wdk

// AlibabawdkbilllistApiResult 结构体
type AlibabawdkbilllistApiResult struct {
	// 响应错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 响应错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 响应返回值
	Model *TxdBillListGetResult `json:"model,omitempty" xml:"model,omitempty"`
	// 响应成功失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
