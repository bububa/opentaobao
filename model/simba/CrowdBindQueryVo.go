package simba

// CrowdBindQueryVo 结构体
type CrowdBindQueryVo struct {
	// 计划id集合
	CampaignIdList []int64 `json:"campaign_id_list,omitempty" xml:"campaign_id_list>int64,omitempty"`
	// 单元id集合,单元已经存在场景必填
	AdgroupIdList []int64 `json:"adgroup_id_list,omitempty" xml:"adgroup_id_list>int64,omitempty"`
	// 偏移量
	Offset int64 `json:"offset,omitempty" xml:"offset,omitempty"`
	// 页面大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 计划id,计划已经存在场景必填
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
}
