package cainiaocntec

// CainiaocnteccompassrpaexeresultsaveResult 结构体
type CainiaocnteccompassrpaexeresultsaveResult struct {
	// 请求trace_id
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 错误code
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 请求结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
