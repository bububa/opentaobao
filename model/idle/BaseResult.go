package idle

// BaseResult 结构体
type BaseResult struct {
	// 是否发货成功
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误描述
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 请求是否执行成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
