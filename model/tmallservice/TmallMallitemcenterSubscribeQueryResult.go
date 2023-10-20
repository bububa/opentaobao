package tmallservice

// TmallmallitemcentersubscribequeryResult 结构体
type TmallmallitemcentersubscribequeryResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 返回结果model
	ResultData *ResultData `json:"result_data,omitempty" xml:"result_data,omitempty"`
	// 错误类型
	ErrorType int64 `json:"error_type,omitempty" xml:"error_type,omitempty"`
	// 是否正常
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
