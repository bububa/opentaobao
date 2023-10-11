package simba

// WordPackageQueryVo 结构体
type WordPackageQueryVo struct {
	// 计划id集合,计划已经存在场景必填
	CampaignIdList []int64 `json:"campaign_id_list,omitempty" xml:"campaign_id_list>int64,omitempty"`
	// 单元id集合,单元已经存在场景必填
	AdgroupIdList []int64 `json:"adgroup_id_list,omitempty" xml:"adgroup_id_list>int64,omitempty"`
}
