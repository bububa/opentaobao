package ihome

// AlibabaIhomeCtomContentPublishStatusApiResult 结构体
type AlibabaIhomeCtomContentPublishStatusApiResult struct {
	// result对象根节点
	Data *StatusResult `json:"data,omitempty" xml:"data,omitempty"`
	// 调用状态
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
}
