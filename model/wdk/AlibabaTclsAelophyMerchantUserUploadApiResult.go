package wdk

// AlibabaTclsAelophyMerchantUserUploadApiResult 结构体
type AlibabaTclsAelophyMerchantUserUploadApiResult struct {
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息描述
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// model
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
}
