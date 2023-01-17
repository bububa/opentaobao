package legalsuit

// StandpointDraftOutPutDto 结构体
type StandpointDraftOutPutDto struct {
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 创建人
	Creator string `json:"creator,omitempty" xml:"creator,omitempty"`
	// 口径表述
	StandpointDesc string `json:"standpoint_desc,omitempty" xml:"standpoint_desc,omitempty"`
	// 建议口径
	DefenseCaliber string `json:"defense_caliber,omitempty" xml:"defense_caliber,omitempty"`
	// id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}
