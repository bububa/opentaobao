package wdk

// AlibabaWdkElemeBillDetailGetApiResult 结构体
type AlibabaWdkElemeBillDetailGetApiResult struct {
	// 错误描述
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 错误编码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 账单信息
	Model *EleBillBo `json:"model,omitempty" xml:"model,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
