package wdk

// AlibabaWdkUmsInventoryPublishApiResult 结构体
type AlibabaWdkUmsInventoryPublishApiResult struct {
	// 调用服务返回结果对象
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 调用服务返回错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 调用服务返回错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 调用服务返回成功失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
