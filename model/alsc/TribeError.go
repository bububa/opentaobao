package alsc

// TribeError 结构体
type TribeError struct {
	// 错误原因
	View string `json:"view,omitempty" xml:"view,omitempty"`
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误原因
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}
