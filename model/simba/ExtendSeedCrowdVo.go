package simba

// ExtendSeedCrowdVo 结构体
type ExtendSeedCrowdVo struct {
	// 人群名称
	CrowdName string `json:"crowd_name,omitempty" xml:"crowd_name,omitempty"`
	// 人群主键id
	CrowdId int64 `json:"crowd_id,omitempty" xml:"crowd_id,omitempty"`
	// 种子人群label信息
	Label *ExtendSeedLabelVo `json:"label,omitempty" xml:"label,omitempty"`
}
