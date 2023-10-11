package simba

// WordPackageSuggestQueryVo 结构体
type WordPackageSuggestQueryVo struct {
	// 类型,overall:综合推荐,trendtheme:趋势词包,automatch:流量智选,extend:流量扩展,highquality:优质卖点词包推荐(用于卖点词包报告),new_default:新建流程默认推荐词包;
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 宝贝id
	MaterialId int64 `json:"material_id,omitempty" xml:"material_id,omitempty"`
	// 计划id
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
	// 单元id
	AdgroupId int64 `json:"adgroup_id,omitempty" xml:"adgroup_id,omitempty"`
}
