package simba

// CampaignQueryVo 结构体
type CampaignQueryVo struct {
	// 筛选项-主体类型
	StatusList []string `json:"status_list,omitempty" xml:"status_list>string,omitempty"`
	// 筛选项-主体类型
	PromotionTypeList []string `json:"promotion_type_list,omitempty" xml:"promotion_type_list>string,omitempty"`
	// 计划id
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
	// 偏移量
	Offset int64 `json:"offset,omitempty" xml:"offset,omitempty"`
	// 页面大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}
