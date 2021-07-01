package idle

// AlibabaIdleRentItemQueryResult 结构体
type AlibabaIdleRentItemQueryResult struct {
	// 返回素材id
	Data *AlibabaIdleRentItemQueryData `json:"data,omitempty" xml:"data,omitempty"`
	// 数据是否可用
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 错误描述
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
}
