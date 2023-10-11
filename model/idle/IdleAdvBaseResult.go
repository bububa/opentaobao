package idle

// IdleAdvBaseResult 结构体
type IdleAdvBaseResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误原因描述
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
}
