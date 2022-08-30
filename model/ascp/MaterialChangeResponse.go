package ascp

// MaterialChangeResponse 结构体
type MaterialChangeResponse struct {
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 处理结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
