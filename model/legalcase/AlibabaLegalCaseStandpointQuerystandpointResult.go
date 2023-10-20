package legalcase

// AlibabaLegalCaseStandpointQuerystandpointResult 结构体
type AlibabaLegalCaseStandpointQuerystandpointResult struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 返回内容
	Content *Page `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
