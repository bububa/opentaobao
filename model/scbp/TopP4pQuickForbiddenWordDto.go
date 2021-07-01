package scbp

// TopP4pQuickForbiddenWordDto 结构体
type TopP4pQuickForbiddenWordDto struct {
	// 屏蔽词
	ForbiddenWord []string `json:"forbidden_word,omitempty" xml:"forbidden_word>string,omitempty"`
	// 1代表新增屏蔽词，2代表删除屏蔽词
	Action int64 `json:"action,omitempty" xml:"action,omitempty"`
	// 定向推广计划ID
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
}
