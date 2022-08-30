package dmp

// CrowdDto 结构体
type CrowdDto struct {
	// 人群有效期
	ValidDate string `json:"valid_date,omitempty" xml:"valid_date,omitempty"`
	// 人群名称
	CrowdName string `json:"crowd_name,omitempty" xml:"crowd_name,omitempty"`
	// 人群创建时间
	Createtime string `json:"createtime,omitempty" xml:"createtime,omitempty"`
	// 人群id
	CrowdId int64 `json:"crowd_id,omitempty" xml:"crowd_id,omitempty"`
	// 人群覆盖人数
	Coverage int64 `json:"coverage,omitempty" xml:"coverage,omitempty"`
}
