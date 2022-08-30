package wdk

// AlibabaAxChannelSkuStatusUpdateApiResult 结构体
type AlibabaAxChannelSkuStatusUpdateApiResult struct {
	// 调用接口返回错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 调用接口返回错误bian ma
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 调用接口返回结果成功失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
