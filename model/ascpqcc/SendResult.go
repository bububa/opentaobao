package ascpqcc

// SendResult 结构体
type SendResult struct {
	// errorMessage
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 业务系统返回响应
	CancelResponse *CancelSampleRelationResponse `json:"cancel_response,omitempty" xml:"cancel_response,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 业务系统返回结果
	UpdateResponse *UpdateSampleResponse `json:"update_response,omitempty" xml:"update_response,omitempty"`
}
