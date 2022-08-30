package tbk

// CrowdDto 结构体
type CrowdDto struct {
	// 人群描述
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 外部人群编码
	ExternalCrowdCode string `json:"external_crowd_code,omitempty" xml:"external_crowd_code,omitempty"`
	// 人群标签名称
	UcrowdName string `json:"ucrowd_name,omitempty" xml:"ucrowd_name,omitempty"`
	// 更新时间
	UpdateTime string `json:"update_time,omitempty" xml:"update_time,omitempty"`
	// 创建时间
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// 状态，0:待生效，1:生效，2:失效
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 人群类型，1:社群
	UcrowdType int64 `json:"ucrowd_type,omitempty" xml:"ucrowd_type,omitempty"`
	// 人群标签id
	UcrowdId int64 `json:"ucrowd_id,omitempty" xml:"ucrowd_id,omitempty"`
	// 覆盖会员数量
	MemberSize int64 `json:"member_size,omitempty" xml:"member_size,omitempty"`
}
