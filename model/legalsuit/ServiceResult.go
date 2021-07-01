package legalsuit

// ServiceResult 结构体
type ServiceResult struct {
	// 案件内容
	Content *CaseModel `json:"content,omitempty" xml:"content,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
}
