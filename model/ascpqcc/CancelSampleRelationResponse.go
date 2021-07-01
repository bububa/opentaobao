package ascpqcc

// CancelSampleRelationResponse 结构体
type CancelSampleRelationResponse struct {
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误code
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 成功标示
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
