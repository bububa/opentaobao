package flight

// AlitripPolicyNormalUploadResult 结构体
type AlitripPolicyNormalUploadResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误原因
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 创建&amp;删除结果参数
	Data *PolicyCreateResponseDto `json:"data,omitempty" xml:"data,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
