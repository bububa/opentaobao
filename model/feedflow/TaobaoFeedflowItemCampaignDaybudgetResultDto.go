package feedflow

// TaobaofeedflowitemcampaigndaybudgetResultDto 结构体
type TaobaofeedflowitemcampaigndaybudgetResultDto struct {
	// 信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 预算总额
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
	// 成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
