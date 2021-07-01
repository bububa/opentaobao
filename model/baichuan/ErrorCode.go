package baichuan

// ErrorCode 结构体
type ErrorCode struct {
	// 详细错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
}
