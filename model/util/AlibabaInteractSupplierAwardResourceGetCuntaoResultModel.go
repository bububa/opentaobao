package util

// AlibabaInteractSupplierAwardResourceGetCuntaoResultModel 结构体
type AlibabaInteractSupplierAwardResourceGetCuntaoResultModel struct {
	// 错误的描述信息（如果有）
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码（如果有）
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 返回值,里面只返回资源ID
	Model *AwardActivityDetailDto `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功，true/false
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
