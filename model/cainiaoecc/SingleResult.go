package cainiaoecc

// SingleResult 结构体
type SingleResult struct {
	// 数据对象
	Result *DelayExceptionCountDto `json:"result,omitempty" xml:"result,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 是否重试
	IsRetry bool `json:"is_retry,omitempty" xml:"is_retry,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 错误描述
	ErrorDesc string `json:"error_desc,omitempty" xml:"error_desc,omitempty"`
}
