package simba

// TopInfo 结构体
type TopInfo struct {
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 是否成功
	Ok bool `json:"ok,omitempty" xml:"ok,omitempty"`
}
