package wdk

// AlibabaWdkElemeBillGetApiResult 结构体
type AlibabaWdkElemeBillGetApiResult struct {
	// 账单信息
	Model *EleBillBo `json:"model,omitempty" xml:"model,omitempty"`
	// 错误描述
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 错误编码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
