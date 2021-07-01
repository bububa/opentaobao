package fenxiao

// TmallSupplychainChannelProductDownshelfResultDto 结构体
type TmallSupplychainChannelProductDownshelfResultDto struct {
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 下架结果
	Module bool `json:"module,omitempty" xml:"module,omitempty"`
	// 错误码
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
}
