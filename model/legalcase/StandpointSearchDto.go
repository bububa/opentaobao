package legalcase

// StandpointSearchDto 结构体
type StandpointSearchDto struct {
	// 扩展属性列表
	Options []Option `json:"options,omitempty" xml:"options>option,omitempty"`
	// 答辩口径
	DefenseCaliber string `json:"defense_caliber,omitempty" xml:"defense_caliber,omitempty"`
	// 生成时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 口径描述
	StandpointDesc string `json:"standpoint_desc,omitempty" xml:"standpoint_desc,omitempty"`
	// 口径id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}
