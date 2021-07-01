package uscesl

// LightResultInfoBo 结构体
type LightResultInfoBo struct {
	// 成功数量
	SuccessCount int64 `json:"success_count,omitempty" xml:"success_count,omitempty"`
	// 失败数量
	FailCount int64 `json:"fail_count,omitempty" xml:"fail_count,omitempty"`
	// 通知消息
	NotifyMessage string `json:"notify_message,omitempty" xml:"notify_message,omitempty"`
}
