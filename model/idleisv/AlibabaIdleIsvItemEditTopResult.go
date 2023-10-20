package idleisv

// AlibabaIdleIsvItemEditTopResult 结构体
type AlibabaIdleIsvItemEditTopResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// data
	Data *IdleItemApiDo `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
