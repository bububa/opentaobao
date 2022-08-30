package dt

// TribeError 结构体
type TribeError struct {
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误说明
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}
