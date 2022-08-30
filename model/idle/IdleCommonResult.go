package idle

// IdleCommonResult 结构体
type IdleCommonResult struct {
	// 错误码
	ErrCodeInfo string `json:"err_code_info,omitempty" xml:"err_code_info,omitempty"`
	// 上下文信息
	Context string `json:"context,omitempty" xml:"context,omitempty"`
	// 错误信息
	ErrMsgInfo string `json:"err_msg_info,omitempty" xml:"err_msg_info,omitempty"`
	// id
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 成功
	RespSuccess bool `json:"resp_success,omitempty" xml:"resp_success,omitempty"`
}
