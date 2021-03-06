package feedflow

// TaobaoFeedflowItemCampaignDeleteResultDto 结构体
type TaobaoFeedflowItemCampaignDeleteResultDto struct {
	// 信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 操作结果
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
