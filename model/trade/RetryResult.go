package trade

// RetryResult 结构体
type RetryResult struct {
	// 扩展参数
	ExtInfo string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
	// 错误信息
	Error *TribeError `json:"error,omitempty" xml:"error,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 是否可重试
	Retry bool `json:"retry,omitempty" xml:"retry,omitempty"`
}
