package alihouse

// ProjectUpdateItemInfoDto 结构体
type ProjectUpdateItemInfoDto struct {
	// 楼盘的外部id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 楼盘描述介绍
	ProjectDescIntroduce string `json:"project_desc_introduce,omitempty" xml:"project_desc_introduce,omitempty"`
}
