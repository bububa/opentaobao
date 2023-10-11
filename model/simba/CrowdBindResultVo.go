package simba

// CrowdBindResultVo 结构体
type CrowdBindResultVo struct {
	// 定向关联关系
	CrowdList []CrowdRefVo `json:"crowd_list,omitempty" xml:"crowd_list>crowd_ref_vo,omitempty"`
	// 计划id,计划已经存在场景必填,eg:添加主体、编辑计划状态等场景
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
	// 单元id
	AdgroupId int64 `json:"adgroup_id,omitempty" xml:"adgroup_id,omitempty"`
}
