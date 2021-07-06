package campus

// BaseResult 结构体
type BaseResult struct {
	// requestId
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误等级
	ErrorLevel string `json:"error_level,omitempty" xml:"error_level,omitempty"`
	// 错误描述
	ErrorExtInfo string `json:"error_ext_info,omitempty" xml:"error_ext_info,omitempty"`
	// 调用信息
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
