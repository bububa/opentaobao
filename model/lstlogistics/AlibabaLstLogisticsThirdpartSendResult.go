package lstlogistics

// AlibabaLstLogisticsThirdpartSendResult 结构体
type AlibabaLstLogisticsThirdpartSendResult struct {
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 错误描述
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 返回内容
	Content *Content `json:"content,omitempty" xml:"content,omitempty"`
}
