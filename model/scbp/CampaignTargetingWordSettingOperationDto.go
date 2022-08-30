package scbp

// CampaignTargetingWordSettingOperationDto 结构体
type CampaignTargetingWordSettingOperationDto struct {
	// 操作优推类型 add-增 del-删 mod-改
	Operation string `json:"operation,omitempty" xml:"operation,omitempty"`
	// 关键词
	Keyword string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// 待操作得优推品id
	AdgroupId string `json:"adgroup_id,omitempty" xml:"adgroup_id,omitempty"`
	// 关键词id
	KeywordId int64 `json:"keyword_id,omitempty" xml:"keyword_id,omitempty"`
	// 计划id
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
}
