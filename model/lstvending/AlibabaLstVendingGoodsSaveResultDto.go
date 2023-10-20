package lstvending

// AlibabalstvendinggoodssaveResultDto 结构体
type AlibabalstvendinggoodssaveResultDto struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 记录唯一值
	Key string `json:"key,omitempty" xml:"key,omitempty"`
}
