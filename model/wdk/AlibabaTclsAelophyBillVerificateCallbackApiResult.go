package wdk

// AlibabaTclsAelophyBillVerificateCallbackApiResult 结构体
type AlibabaTclsAelophyBillVerificateCallbackApiResult struct {
	// 回调是否处理成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 错误说明
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 错误编码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
}
