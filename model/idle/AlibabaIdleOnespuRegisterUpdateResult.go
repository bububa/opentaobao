package idle

// AlibabaIdleOnespuRegisterUpdateResult 结构体
type AlibabaIdleOnespuRegisterUpdateResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 返回值
	Module string `json:"module,omitempty" xml:"module,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
