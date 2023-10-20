package mtopopen

// AlibabainteractaopdataregisterResult 结构体
type AlibabainteractaopdataregisterResult struct {
	// xx
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// xxx失败
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// xxx失败
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 跟踪请求使用
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 接口调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
