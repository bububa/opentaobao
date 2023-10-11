package miniapp

// MiniappResult 结构体
type MiniappResult struct {
	// model
	Model []AppChannelConfigDto `json:"model,omitempty" xml:"model>app_channel_config_dto,omitempty"`
	// 错误码
	ErrorType string `json:"error_type,omitempty" xml:"error_type,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrType int64 `json:"err_type,omitempty" xml:"err_type,omitempty"`
	// true or false
	Successful bool `json:"successful,omitempty" xml:"successful,omitempty"`
}
