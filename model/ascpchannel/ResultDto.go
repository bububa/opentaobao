package ascpchannel

// ResultDto 结构体
type ResultDto struct {
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 请求是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 是否绑定成功
	Module bool `json:"module,omitempty" xml:"module,omitempty"`
}
