package wdk

// AlibabaWdkAxStoreQueryApiResult 结构体
type AlibabaWdkAxStoreQueryApiResult struct {
	// 调用接口返回错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 调用接口返回错误编码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 查询接口返回对象
	Model *AxStoreQueryResponse `json:"model,omitempty" xml:"model,omitempty"`
	// 调用接口返回成功失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
