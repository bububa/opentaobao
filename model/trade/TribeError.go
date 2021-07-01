package trade

// TribeError 结构体
type TribeError struct {
	// 错误可读性描述
	View string `json:"view,omitempty" xml:"view,omitempty"`
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}
