package lstvending

// AlibabalstvendingcargospacesaveResultDto 结构体
type AlibabalstvendingcargospacesaveResultDto struct {
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 记录唯一标识
	Key string `json:"key,omitempty" xml:"key,omitempty"`
}
