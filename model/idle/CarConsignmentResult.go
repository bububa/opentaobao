package idle

// CarConsignmentResult 结构体
type CarConsignmentResult struct {
	// 错误信息
	PerformError string `json:"perform_error,omitempty" xml:"perform_error,omitempty"`
	// 请求成功
	TransformSuccess bool `json:"transform_success,omitempty" xml:"transform_success,omitempty"`
}
