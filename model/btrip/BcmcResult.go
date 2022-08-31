package btrip

// BcmcResult 结构体
type BcmcResult struct {
	// 错误信息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 错误码
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 成功标识
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
