package hotelhstdf

// DidaResult 结构体
type DidaResult struct {
	// 错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 错误信息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 服务是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
