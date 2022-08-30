package logistic

// IsvResult 结构体
type IsvResult struct {
	// 共享码
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 错误码
	ServerErrorCode string `json:"server_error_code,omitempty" xml:"server_error_code,omitempty"`
	// 描述
	Describe string `json:"describe,omitempty" xml:"describe,omitempty"`
}
