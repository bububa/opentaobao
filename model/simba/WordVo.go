package simba

// WordVo 结构体
type WordVo struct {
	// 词
	Word string `json:"word,omitempty" xml:"word,omitempty"`
	// 关键词基础出价
	BidPrice string `json:"bid_price,omitempty" xml:"bid_price,omitempty"`
	// 状态,reject:排查拒绝,preview:排查中,launch:投放中,noflow:无展现
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 计划id,计划已经存在场景必填,eg:添加主体、编辑计划状态等场景
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
	// 单元id
	AdgroupId int64 `json:"adgroup_id,omitempty" xml:"adgroup_id,omitempty"`
	// 匹配类型,1:精准匹配,4:广泛匹配
	MatchScope int64 `json:"match_scope,omitempty" xml:"match_scope,omitempty"`
	// 词id
	BidwordId int64 `json:"bidword_id,omitempty" xml:"bidword_id,omitempty"`
	// 抢位信息
	BidStrategyInfo *BidStrategyVo `json:"bid_strategy_info,omitempty" xml:"bid_strategy_info,omitempty"`
}
