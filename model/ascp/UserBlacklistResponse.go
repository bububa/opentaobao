package ascp

// UserBlacklistResponse 结构体
type UserBlacklistResponse struct {
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 系统成功失败   true|false
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 是否可重试
	IsRetry bool `json:"is_retry,omitempty" xml:"is_retry,omitempty"`
}
