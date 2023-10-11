package simba

// CampaignBrandFilterVo 结构体
type CampaignBrandFilterVo struct {
	// deeplink维度过滤,Discover:Discover 发现,Engage:Engage 种草,Enthuse:Enthuse 互动,Perform:Perform 行动,Initial:Initial 首购,Numerous:Numerous 复购,Keen:Keen 至爱
	DeeplinkList []string `json:"deeplink_list,omitempty" xml:"deeplink_list>string,omitempty"`
	// 品牌id
	BrandId string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
}
