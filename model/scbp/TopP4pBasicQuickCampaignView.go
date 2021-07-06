package scbp

// TopP4pBasicQuickCampaignView 结构体
type TopP4pBasicQuickCampaignView struct {
	// 屏蔽词列表
	ForbiddenWords []string `json:"forbidden_words,omitempty" xml:"forbidden_words>string,omitempty"`
	// 出价区间-上限
	MaxPrice string `json:"max_price,omitempty" xml:"max_price,omitempty"`
	// 出价区间-下限
	MinPrice string `json:"min_price,omitempty" xml:"min_price,omitempty"`
	// 计划标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 计划每日预算
	Budget int64 `json:"budget,omitempty" xml:"budget,omitempty"`
	// 推广商品数量
	ProductCount int64 `json:"product_count,omitempty" xml:"product_count,omitempty"`
	// 计划状态(0暂停、1推广中、-1点爆)
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 定向推广计划ID
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
}
