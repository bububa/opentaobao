package examination

// ReportImTokenStatusResponse 结构体
type ReportImTokenStatusResponse struct {
	// 令牌校验失败原因
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 令牌校验状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
}
