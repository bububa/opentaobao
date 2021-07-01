package ascpffo

// AliexpressFulfillmentEventResult 结构体
type AliexpressFulfillmentEventResult struct {
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 错误码
	ErrorCode *ErrorCode `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 是否可重试
	Retry bool `json:"retry,omitempty" xml:"retry,omitempty"`
}
