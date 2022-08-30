package alsc

// SingleDataResult 结构体
type SingleDataResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 结果码
	ResultStatus int64 `json:"result_status,omitempty" xml:"result_status,omitempty"`
	// 返回数据
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}
