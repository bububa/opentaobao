package paimai

// Result4Top 结构体
type Result4Top struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 处理结果
	Value bool `json:"value,omitempty" xml:"value,omitempty"`
}
