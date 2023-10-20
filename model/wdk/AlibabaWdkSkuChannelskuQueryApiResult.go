package wdk

// AlibabaWdkSkuChannelskuQueryApiResult 结构体
type AlibabaWdkSkuChannelskuQueryApiResult struct {
	// 异常状态码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 异常信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 业务数据模型
	Model *ChannelSkuDo `json:"model,omitempty" xml:"model,omitempty"`
	// 业务调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
